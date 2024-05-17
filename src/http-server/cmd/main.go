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
	config, err := config.LoadServerConfigFromFile(*configFile)
	if err != nil {
		logrus.Fatal("load config error:", err)
		return
	}
	configBytes, err := json.Marshal(config)
	log.Info("config:", string(configBytes))

	// 不晓得什么作用
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

	// 添加grpc, net listen 监听grpcEndpoint的连接
	l, err := net.Listen("tcp", grpcEndpoint)
	if err != nil {
		log.Fatal("failed to listen :", err)
	}
	log.Info("gRPC listen :", grpcEndpoint)
	go func() {
		if err := grpcServer.Serve(l); err != nil {
			log.Warn("gRPC server closed: ", err)
		}
	}()
	dbService := service.DbService{}
	db.RegisterDBApiServer(grpcServer, &dbService)
	//mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{OrigName: true, EmitDefaults: true}))

	dialOptions := []grpc.DialOption{grpc.WithInsecure(), grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(16 * MB))}

	gmux := runtime.NewServeMux(runtime.WithMarshalerOption(
		runtime.MIMEWildcard, &runtime.JSONPb{OrigName: true, EmitDefaults: true}),
		runtime.WithMarshalerOption("application/smalljson", &runtime.JSONPb{OrigName: true, EmitDefaults: false}),
	)
	if err := db.RegisterDBApiHandlerFromEndpoint(ctx, gmux, ":10608", dialOptions); err != nil {
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
