package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gogo/protobuf/proto"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"http-server/api/db"
	config "http-server/config"
	"http-server/dao"
	"http-server/service"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const (
	_ = iota // ignore first value by assigning to blank identifier
	// KB : kilobyte
	KB = 1 << (10 * iota)
	// MB : megabyte
	MB
)

var (
	// 从入参中获取config参数 默认值为config/server.json
	configFile = flag.String("config", "config/server.json", "server config file path")
	grpcServer *grpc.Server
	httpServer *http.Server
)

func main() {
	// 解析入参参数
	flag.Parse()
	fmt.Println("http-server test")
	conf, err := config.LoadServerConfigFromFile(*configFile)
	if err != nil {
		logrus.Fatal("load config error:", err)
		return
	}
	configBytes, err := json.Marshal(conf)
	log.Info("config:", string(configBytes))

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	grpcLog := log.WithField("server", "entity")

	opt := grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		grpc_logrus.UnaryServerInterceptor(grpcLog),
		grpc_prometheus.UnaryServerInterceptor,
		payloadUnaryServerInterceptor(),
	))
	grpcEndpoint := ":10608"
	httpEndpoint := ":10607"
	grpcServer = grpc.NewServer(opt)
	databaseClient, err := config.NewDatabaseClient(conf)
	if err != nil {
		log.Error("config init error.", err)
		return
	}
	dbQueryDao := dao.DBQuery{Db: databaseClient.Db}
	dbService := service.DbService{DbQueryDao: dbQueryDao}
	// gRPC 服务器会处理来自客户端的 gRPC 请求，并调用 DbService 中相应的方法来处理这些请求
	db.RegisterDBApiServer(grpcServer, &dbService)
	// 创建了一个 gRPC 服务器，并在指定的地址和端口上开始监听，同时使用一个单独的 Go 协程处理连接
	// 使用 net.Listen 创建一个 TCP 的监听器，指定了地址和端口
	l, err := net.Listen("tcp", grpcEndpoint)
	if err != nil {
		log.Fatal("failed to listen :", err)
	}
	log.Info("gRPC listen :", grpcEndpoint)
	go func() {
		// 开始监听来自 l 的连接，并处理 gRPC 请求。如果出现错误，通过日志记录错误信息。
		if err := grpcServer.Serve(l); err != nil {
			log.Warn("gRPC server closed: ", err)
		}
	}()

	//mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{OrigName: true, EmitDefaults: true}))

	dialOptions := []grpc.DialOption{grpc.WithInsecure(), grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(16 * MB))}

	//创建了一个新的  gRPC Gateway 的 ServeMux，它将 gRPC 服务映射到 RESTful API。在这里，配置了两种 Marshaler，一种是默认的 JSON 格式，另一种是用于自定义 MIME 类型 "application/smalljson" 的 JSON 格式
	gmux := runtime.NewServeMux(runtime.WithMarshalerOption(
		runtime.MIMEWildcard, &runtime.JSONPb{OrigName: true, EmitDefaults: true}),
		runtime.WithMarshalerOption("application/smalljson", &runtime.JSONPb{OrigName: true, EmitDefaults: false}),
	)

	// 将gRPC服务的处理程序 注册到 gRPC Gateway中 gRPC Gateway 就能够将来自客户端的 HTTP 请求转发到相应的 gRPC 服务中处理。
	if err := db.RegisterDBApiHandlerFromEndpoint(ctx, gmux, grpcEndpoint, dialOptions); err != nil {
		log.Fatal("Failed to register service handler for grpc: ", err)
	}

	httpServer = &http.Server{Addr: httpEndpoint, Handler: gmux}
	log.Info("start listening Http proxy at ", httpEndpoint)
	go func() {
		err := httpServer.ListenAndServe()
		if err != nil {
			log.Warn("http server stopped:", err)
		}
		log.Info("http listening at ", httpEndpoint)
	}()
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	logrus.Info("exiting...")
}

func payloadUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if log.GetLevel() >= log.DebugLevel {
			log.Debug("grpc.request:", req.(proto.Message).String())
		}
		resp, err := handler(ctx, req)
		return resp, err
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}
