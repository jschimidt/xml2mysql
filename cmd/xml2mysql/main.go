package main
import ("flag";"fmt";"os";"github.com/jschimidt/xml2mysql/internal/analyzer";"github.com/jschimidt/xml2mysql/internal/config")
const version="0.1.0"
func main(){if len(os.Args)<2{usage();os.Exit(2)};switch os.Args[1]{case "analyze":fs:=flag.NewFlagSet("analyze",flag.ExitOnError);p:=fs.String("config","/app/config.yaml","configuration YAML");w:=fs.Int("workers",0,"override workers");_=fs.Parse(os.Args[2:]);c,e:=config.Load(*p);if e!=nil{fatal(e)};if *w>0{c.Analysis.Workers=*w};if e:=analyzer.Run(c,version);e!=nil{fatal(e)};case "version","--version","-v":fmt.Println(version);default:usage();os.Exit(2)}}
func usage(){fmt.Fprintln(os.Stderr,"xml2mysql v"+version);fmt.Fprintln(os.Stderr,"usage: xml2mysql analyze --config /app/config.yaml [--workers N]")}
func fatal(e error){fmt.Fprintln(os.Stderr,"error:",e);os.Exit(1)}
