package main

import (
    "log"
    "fmt"
    "net/http"
    "os"
    "strconv"
    "github.com/toresbe/caspar_go"
    "net"
)

type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"message": "hello world"}`))
}

func listen_port(default_port uint16) (configured_port uint16) {
    port_env_str, port_env_set := os.LookupEnv("PORT")

    if port_env_set {
        port_env_uint, conv_error := strconv.ParseUint(port_env_str, 10, 16)
        if conv_error == nil {
            configured_port = uint16(port_env_uint)
            log.Printf("Configured to listen on port %v", configured_port)
            return
        } else {
            log.Printf("PORT environment has invalid value \"%v\", defaulting to %v", port_env_str, default_port)
        }
    } else {
        log.Printf("PORT environment not set, listening on %v", default_port)
    }

    return default_port
}

func main() {
    s := &server{}
    http.Handle("/", s)
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", listen_port(1234)), nil))
    dummy_conn, _ := net.Pipe()
    caspar_go.Open(dummy_conn)
}
