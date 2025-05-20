// comprobadns
// https://github.com/drieiro/comprobadns

package main

import (
	"fmt"
	"net"
	"os"
	"text/tabwriter"
)

const (
	programName    = "comprobadns"
	programVersion = "1.3.2"
)

func main() {
	fmt.Printf("%s v%s\n\n", programName, programVersion)

	// Comprobar se se pasou dominio por argumento
	if len(os.Args) > 1 {
		domain := os.Args[1]
		checkDNS(domain)
		return
	}

	// Modo interactivo
	for {
		var domain string
		fmt.Print("Introduce un dominio: ")
		if _, err := fmt.Scanln(&domain); err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler o dominio: %v\n", err)
			return
		}
		checkDNS(domain)
		fmt.Println()
	}
}

func checkDNS(domain string) {
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)

	// Rexistro @
	if ips, err := net.LookupIP(domain); err == nil {
		for _, ip := range ips {
			ptr := lookupPTR(ip)
			fmt.Fprintf(w, "@\t%s --> %v\n", ip, ptr)
		}
	} else {
		fmt.Fprintf(os.Stderr, "Non foi posible obter o rexistro A (@): %v\n", err)
	}

	// Rexistro WWW
	www := "www." + domain
	if ips, err := net.LookupIP(www); err == nil {
		for _, ip := range ips {
			ptr := lookupPTR(ip)
			fmt.Fprintf(w, "WWW\t%s --> %v\n", ip, ptr)
		}
	} else {
		fmt.Fprintf(os.Stderr, "Non foi posible obter o rexistro A (www): %v\n", err)
	}

	// Rexistros MX
	if mxs, err := net.LookupMX(domain); err == nil {
		for _, mx := range mxs {
			ips := lookupIPOrPlaceholder(mx.Host)
			fmt.Fprintf(w, "MX\t%s --> %v\n", mx.Host, ips)
		}
	} else {
		fmt.Fprintf(os.Stderr, "Non foi posible obter o rexistro MX: %v\n", err)
	}

	// Rexistros NS
	if nss, err := net.LookupNS(domain); err == nil {
		for _, ns := range nss {
			ips := lookupIPOrPlaceholder(ns.Host)
			fmt.Fprintf(w, "NS\t%s --> %v\n", ns.Host, ips)
		}
	} else {
		fmt.Fprintf(os.Stderr, "Non foi posible obter o rexistro NS: %v\n", err)
	}

	// Rexistros TXT
	if txts, err := net.LookupTXT(domain); err == nil {
		for _, txt := range txts {
			fmt.Fprintf(w, "TXT\t%s\n", txt)
		}
	} else {
		fmt.Fprintf(os.Stderr, "Non foi posible obter o rexistro TXT: %v\n", err)
	}

	// DKIM
	dkim := "default._domainkey." + domain
	if dkimTXT, err := net.LookupTXT(dkim); err == nil {
		for _, val := range dkimTXT {
			fmt.Fprintf(w, "DKIM\t%s\n", val)
		}
	} else {
		fmt.Fprintf(os.Stderr, "Non foi posible obter o rexistro DKIM: %v\n", err)
	}

	// DMARC
	dmarc := "_dmarc." + domain
	if dmarcTXT, err := net.LookupTXT(dmarc); err == nil {
		for _, val := range dmarcTXT {
			fmt.Fprintf(w, "DMARC\t%s\n", val)
		}
	} else {
		fmt.Fprintf(os.Stderr, "Non foi posible obter o rexistro DMARC: %v\n", err)
	}

	w.Flush()
}

// Devolve os rexistros PTR ou "Sen PTR"
func lookupPTR(ip net.IP) []string {
	ptr, err := net.LookupAddr(ip.String())
	if err != nil || len(ptr) == 0 {
		return []string{"Sen PTR"}
	}
	return ptr
}

// Intenta resolver IP, ou devolve placeholder se falla
func lookupIPOrPlaceholder(host string) []net.IP {
	ips, err := net.LookupIP(host)
	if err != nil || len(ips) == 0 {
		return []net.IP{net.IPv4(0, 0, 0, 0)} // Placeholder: 0.0.0.0
	}
	return ips
}
