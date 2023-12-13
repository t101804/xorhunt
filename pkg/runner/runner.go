package runner

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"regexp"

	fileutil "github.com/projectdiscovery/utils/file"
	urlutil "github.com/projectdiscovery/utils/url"
	"github.com/t101804/xorhunt/options"
	"github.com/t101804/xorhunt/pkg/logger"
	"github.com/t101804/xorhunt/pkg/resolve"
	"github.com/t101804/xorhunt/pkg/templateparser"
)

type Runner struct {
	options        *options.GlobalOptions
	resolverClient *resolve.Resolver
	tmpl           *templateparser.Tmpl
}

func NewRunner(options *options.GlobalOptions) (*Runner, error) {
	runner := &Runner{options: options}

	if fileutil.FileExists(runner.options.TemplateName) {
		logger.Info().Msgf("Loading provider config from %s", runner.options.TemplateName)
		runner.loadProvidersFrom(options.TemplateName)
	} else {
		runner.options.TemplateName = "templates/autorecon.yaml"
		logger.Info().Msgf("Loading provider config from the default location: %s", runner.options.TemplateName)
		runner.loadProvidersFrom(runner.options.TemplateName)
	}

	logger.Info().Msg(runner.tmpl.TmplStruct.Config.Site)

	return runner, nil
}

func (r *Runner) Start() error {
	return r.RunWithCtx(context.Background())
}

func (r *Runner) RunWithCtx(ctx context.Context) error {
	outputs := []io.Writer{r.options.Output}

	if r.options.IterFile != "" && fileutil.FileExists(r.options.IterFile) {
		f, err := os.Open(r.options.IterFile)
		if err != nil {
			return err
		}
		err = r.EnumerateListWithCtx(ctx, f, outputs)
		f.Close()
		return err

	}
	return nil
}

func (r *Runner) parseURL(url string) (*urlutil.URL, error) {
	urlx, err := urlutil.ParseURL(url, false)
	if err != nil {
		logger.Debug().Msgf("failed to parse url %v got %v in unsafe:%v", url, err, false)
	}
	return urlx, err
}
func (r *Runner) EnumerateListWithCtx(ctx context.Context, reader io.Reader, writers []io.Writer) error {
	scanner := bufio.NewScanner(reader)
	ip, _ := regexp.Compile(`^([0-9\.]+$)`)
	for scanner.Scan() {
		fmt.Println(">" + scanner.Text())
		domain, err := sanitize(scanner.Text())
		isIp := ip.MatchString(domain)
		if errors.Is(err, errors.New("empty data")) || (r.options.ExcludeIps && isIp) {
			continue
		}
		if domain != "" && r.options.ChangeToIp {
			URL, err := r.parseURL(domain)
			if err != nil {
				return err
			}
			ips, _ := net.LookupIP(URL.Host)
			fmt.Println(ips)
			// for _, ip := range ips {
			// 	if ipv4 := ip.To4(); ipv4 != nil {
			// 		fmt.Println("IPv4: ", ipv4)
			// 	}
			// }
		}
	}
	return nil
}
