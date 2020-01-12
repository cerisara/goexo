
func main() {
        var bindAddress = "localhost"
        m := &autocert.Manager{
                Prompt: autocert.AcceptTOS,
                Cache:  autocert.DirCache("."),
        }
        host, err := url.Parse(app.cfg.App.Host)

        if err != nil {
                log.Error("[WARNING] Unable to parse configured host! %s", err)
                log.Error(`[WARNING] ALL hosts are allowed, which can open you to an attack where
clients connect to a server by IP address and pretend to be asking for an
incorrect host name, and cause you to reach the CA's rate limit for certificate
requests. We recommend supplying a valid host name.`)
                log.Info("Using autocert on ANY host")
        } else {
                log.Info("Using autocert on host %s", host.Host)
                m.HostPolicy = autocert.HostWhitelist(host.Host)
        }
        s := &http.Server{
                Addr:    ":https",
                Handler: r,
                TLSConfig: &tls.Config{
                        GetCertificate: m.GetCertificate,
                },

}

