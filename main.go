// chkdns
// https://github.com/drieiro/chkdns

package main

import (
    "fmt"
    "net"
    "os"
    "regexp"
    "text/tabwriter"
)

const (
    programName = "chkdns"
    programVersion = "1.4"
)

func validateDomainName(domain string) bool {
     RegExp := regexp.MustCompile(`^(([a-zA-Z]{1})|([a-zA-Z]{1}[a-zA-Z]{1})|([a-zA-Z]{1}[0-9]{1})|([0-9]{1}[a-zA-Z]{1})|([a-zA-Z0-9][a-zA-Z0-9-_]{1,61}[a-zA-Z0-9]))\.([a-zA-Z]{2,6}|[a-zA-Z0-9-]{2,30}\.[a-zA-Z]{2,3})$`)
     return RegExp.MatchString(domain)
 }

func main() {
    for {
        fmt.Print("Introduce un dominio: ")
        var domain string
        // Manexo dos erros
            _, err := fmt.Scanln(&domain)
            if err != nil {
                fmt.Fprintln(os.Stderr, err)
                return
            }

        // Validación da variable domain
        if !validateDomainName(domain) {
            fmt.Printf("O valor %s non é un dominio válido.\n", domain)
            return
        } else {
            // Obter rexistros @
            atRecords, err := net.LookupIP(domain)
                if err != nil {
                        fmt.Fprintf(os.Stderr, "Non foi posible obter o rexistro A (@): %v\n", err)
                }

            // Obter rexistros www
            wwwdomain := "www." + domain
            wwwRecords, err := net.LookupIP(wwwdomain)
                if err != nil {
                        fmt.Fprintf(os.Stderr, "Non foi posible obter o rexistro A (www): %v\n", err)
                }

            // Obter rexistros MX
            mxRecords, err := net.LookupMX(domain)
                if err != nil {
                        fmt.Fprintf(os.Stderr, "Non foi posible obter o rexistro MX: %v\n", err)
                }

            // Obter rexistros NS
            nsRecords, err := net.LookupNS(domain)
                if err != nil {
                    fmt.Fprintf(os.Stderr, "Non foi posible obter o rexistro NS: %v\n", err)
                }

            // Obter rexistros TXT
            txtRecords, err := net.LookupTXT(domain)
                if err != nil {
                    fmt.Fprintf(os.Stderr, "Non foi posible obter o rexistro TXT: %v\n", err)
                }

            // Obter rexistro DKIM
            defaultdkim := "default._domainkey." + domain
            dkimRecords, err := net.LookupTXT(defaultdkim)
                if err != nil {
                    fmt.Fprintf(os.Stderr, "Non foi posible obter o rexistro DKIM: %v\n", err)
                }

            // Obter rexistro DMARC
            defaultdmarc := "_dmarc." + domain
            dmarcRecords, err := net.LookupTXT(defaultdmarc)
                if err != nil {
                    fmt.Fprintf(os.Stderr, "Non foi posible obter o rexistro DMARC: %v\n", err)
                }

            fmt.Print("\n")

            // Configuración da táboa
            w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)

            // Mostrar rexistros @
            for _, at := range atRecords {
                // Isto mostra os rexistros PTR
                ptr, err := net.LookupAddr(at.String())
                    if err != nil {
                            ptr = append(ptr, "Sen PTR")
                    }
                fmt.Fprintln(w, "@\t", at, "-->", ptr)
            }

            // Mostrar rexistros WWW
            for _, www := range wwwRecords {
                ptr, err := net.LookupAddr(www.String())
                    if err != nil {
                            ptr = append(ptr, "Sen PTR")
                    }
                fmt.Fprintln(w, "WWW\t", www, "-->", ptr)
            }

            // Mostrar rexistros MX
            for _, mx := range mxRecords {
                // Obter tamén a resolución IP do MX
                mxIP, err := net.LookupIP(mx.Host)
                    if err != nil {
                        mxIP = append(mxIP, net.IP(fmt.Sprintf("Sen IP")))
                    }
                fmt.Fprintln(w, "MX\t", mx.Host, "-->", mxIP)
            }
            // Mostrar rexistros NS
            for _, ns := range nsRecords {
                // Obter tamén a resolución IP dos NS
                nsIP, err := net.LookupIP(ns.Host)
                    if err != nil {
                        nsIP = append(nsIP, net.IP(fmt.Sprintf("Sen IP")))
                    }
                fmt.Fprintln(w, "NS\t", ns.Host, "-->", nsIP)
            }

            // Mostrar rexistros TXT
            for _, txt := range txtRecords {
                fmt.Fprintln(w, "TXT\t", txt)
            }

            // Mostrar o rexistro DKIM principal (default._domainkey)
            for _, dkim := range dkimRecords {
                fmt.Fprintln(w, "DKIM\t", dkim)
            }

            // Mostrar o rexistro DMARC principal (_dmarc)
            for _, dmarc := range dmarcRecords {
                fmt.Fprintln(w, "DMARC\t", dmarc)
            }

            w.Flush()
            fmt.Print("\n")
        }
    }
}
