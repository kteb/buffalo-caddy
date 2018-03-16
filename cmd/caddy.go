package cmd

import (
	"os"
	"os/exec"

	"github.com/gobuffalo/makr"
	"github.com/spf13/cobra"
)

// caddyGenCmd represents the caddy gen command
var caddyGenCmd = &cobra.Command{
	Use:     "caddyfile",
	Aliases: []string{"c"},
	Short:   "Generates a Caddyfile for the local environement",
	Run: func(cmd *cobra.Command, args []string) {
		s := `localhost:80 {
  redir https://localhost
}

localhost:443 {
  tls self_signed
  proxy / http://127.0.0.1:3000/ {
                  transparent
                  websocket
          }
}`

		// create a caddy file
		g := makr.New()
		g.Add(makr.NewFile("Caddyfile", s))
		err := g.Run(".", makr.Data{})
		if err != nil {
			os.Exit(1)
		}
	},
}

// caddyGenCmd represents the caddy gen command
var caddyDevCmd = &cobra.Command{
	Use:     "caddy",
	Aliases: []string{"c"},
	Short:   "Start caddy and buffalo dev",
	Run: func(cmd *cobra.Command, args []string) {
		c := makr.New()
		b := makr.New()

		c.Add(makr.NewCommand(exec.Command("caddy")))
		b.Add(makr.NewCommand(exec.Command("buffalo", "dev")))

		go run(c, makr.Data{})

		run(b, makr.Data{})
	},
}

func run(g *makr.Generator, d makr.Data) {
	err := g.Run(".", d)
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	RootCmd.AddCommand(caddyGenCmd)
	RootCmd.AddCommand(caddyDevCmd)
}
