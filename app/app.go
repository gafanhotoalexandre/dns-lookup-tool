package app

import (
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/urfave/cli"
)

// Gerar vai retornar a command-line pronta para execução
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "DNS Lookup Tool"
	app.Usage = "Ferramenta para consultas DNS - busca IPs e servidores de nomes"
	app.Version = "1.0.0"
	app.Author = "alemartins.dev.br"
	app.Email = "alexandrevmartinsdelima@gmail.com"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "alemartins.dev.br",
			Usage: "Nome do host/domínio para consulta",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "ip",
			Aliases: []string{"i"},
			Usage:   "Busca endereços IP de um domínio",
			Flags:   flags,
			Action:  buscarIps,
		},
		{
			Name:    "servidores",
			Aliases: []string{"ns", "s"},
			Usage:   "Busca servidores de nomes (NS) de um domínio",
			Flags:   flags,
			Action:  buscarServidores,
		},
	}

	return app
}

func buscarIps(c *cli.Context) error {
	host := c.String("host")

	if host == "" {
		fmt.Fprint(os.Stderr, "Erro: host não pode estar vazio\n")
		return cli.NewExitError("host não especificado", 1)
	}

	fmt.Printf("Buscando IPs para: %s\n", host)
	fmt.Println(strings.Repeat("-", 40))

	ips, erro := net.LookupIP(host)
	if erro != nil {
		fmt.Fprintf(os.Stderr, "Erro ao buscar IPs para %s: %v\n", host, erro)
		return cli.NewExitError("falha na consulta DNS", 1)
	}

	for i, ip := range ips {
		tipoIP := "IPv4"
		if ip.To4() == nil {
			tipoIP = "IPv6"
		}
		fmt.Printf("%d. %s (%s)\n", i+1, ip, tipoIP)
	}

	return nil
}

func buscarServidores(c *cli.Context) error {
	host := c.String("host")

	if host == "" {
		fmt.Fprintf(os.Stderr, "Erro: host não pode estar vazio\n")
		return cli.NewExitError("host não especificado", 1)
	}

	fmt.Printf("Buscando servidores NS para: %s\n", host)
	fmt.Println(strings.Repeat("-", 40))

	servidores, erro := net.LookupNS(host)
	if erro != nil {
		fmt.Fprintf(os.Stderr, "Erro ao buscar servidores para %s: %v\n", host, erro)
		return cli.NewExitError("falha na consulta DNS", 1)
	}

	if len(servidores) == 0 {
		fmt.Println("Nenhum servidor NS encontrado")
		return nil
	}

	for i, servidor := range servidores {
		fmt.Printf("%d. %s\n", i+1, servidor.Host)
	}

	return nil
}
