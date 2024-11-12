package main

// BY MALTASH ON GITHUB

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Magenta = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

func main() {
	siteURL := flag.String("u", "", "Website URL")
	flag.Parse()

	if *siteURL == "" {
		fmt.Println("Please specify the website (https://example.com) using the -u argument")
		os.Exit(1)
	}
	fmt.Println(" ██████╗  ██████╗ ██╗    ██╗██████╗ ")
	fmt.Println("██╔════╝ ██╔═══██╗██║    ██║██╔══██╗")
	fmt.Println("██║  ███╗██║   ██║██║ █╗ ██║██████╔╝")
	fmt.Println("██║   ██║██║   ██║██║███╗██║██╔═══╝ ")
	fmt.Println("╚██████╔╝╚██████╔╝╚███╔███╔╝██║     ")
	fmt.Println(" ╚═════╝  ╚═════╝  ╚══╝╚══╝ ╚═╝     ")
	fmt.Println("")
	fmt.Println("GOWP v1 - A lightweight WordPress site scanner.")
	fmt.Print(Blue + "By MALTASH | " + Reset)
	fmt.Println(Red + " I am not responsible for your actions, this script was for educational purposes only." + Reset)
	fmt.Println("")
	time.Sleep(2 * time.Second)
	fmt.Println(Yellow + "[INFO] The scan is in progress... Please wait." + Reset)
	fmt.Println("")
	pingWebsite(*siteURL)
	testWpAdmin(*siteURL)
	testUpload(*siteURL)
	testThemeLoc(*siteURL)
	testPlugin(*siteURL)
	fmt.Println(Yellow + "[INFO] Scan for sitemaps in progress..." + Reset)
	fmt.Println("")
	testSitemap(*siteURL)
	fmt.Println(Yellow + "[INFO] Scan for interesting files in progress..." + Reset)
	fmt.Println("")
	testRobotxt(*siteURL)
	testXMLRPC(*siteURL)
}

func pingWebsite(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(Red + "[-] The website is not accessible, scan canceled." + Reset)
		os.Exit(1)
	}

	defer resp.Body.Close()

	fmt.Println(Green + "[+] The website is online" + Reset)
	fmt.Println("-> Found with a ping")
	fmt.Println("")
}

func testWpAdmin(url string) {
	adminURL := url + "/wp-admin"
	resp, err := http.Get(adminURL)
	if err != nil {
		fmt.Println(Red + "Error accessing wp-admin" + Reset)
		fmt.Println("")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound || resp.StatusCode == http.StatusForbidden {
		fmt.Println(Red + "[-] wp-admin not accessible" + Reset)
		fmt.Println("")
	} else {
		fmt.Println(Green + "[+] wp-admin accessible" + Reset)
		fmt.Println("-> Location: " + url + "wp-admin")
	}
}

func testPlugin(url string) {
	pluginURL := url + "/wp-content/plugins/"
	resp, err := http.Get(pluginURL)
	if err != nil {
		fmt.Println(Red + "[-] Plugins directory not available" + Reset)
		fmt.Println("")
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotFound || resp.StatusCode == http.StatusForbidden {
		fmt.Println(Red + "[-] Plugins directory not accessible" + Reset)
		fmt.Println("")
	} else {
		fmt.Println(Green + "[+] Plugins directory found" + Reset)
		fmt.Println("-> Found through a direct request")
		fmt.Println("-> Location: " + url + "wp-content/plugins/")
		fmt.Println("")
	}
}

func testUpload(url string) {
	uploadURL := url + "/wp-content/uploads/"
	resp, err := http.Get(uploadURL)
	if err != nil {
		fmt.Println(Red + "[-] Upload directory not available" + Reset)
		fmt.Println("")
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound || resp.StatusCode == http.StatusForbidden {
		fmt.Println(Red + "[-] Upload directory not accessible" + Reset)
		fmt.Println("")
	} else {
		fmt.Println(Green + "[+] Upload directory found" + Reset)
		fmt.Println("-> Found through a direct request")
		fmt.Println("-> Location: " + url + "wp-content/uploads/")
		fmt.Println("")
	}
}

func testSitemap(url string) {
	testURLs := []string{
		"/sitemap_index.xml",
		"/wp-sitemap-posts-post-1.xml",
		"/wp-sitemap-posts-post-2.xml",
		"/wp-sitemap-posts-page-1.xml",
	}

	for _, path := range testURLs {
		fullURL := url + path
		resp, err := http.Get(fullURL)
		if err != nil {
			fmt.Printf(Red+"[-] %s not accessible: %s\n\n"+Reset, path, err)
			continue
		}

		defer resp.Body.Close()

		if resp.StatusCode == http.StatusNotFound || resp.StatusCode == http.StatusForbidden {
			fmt.Printf(Red+"[-] %s not accessible: %s\n\n"+Reset, path, resp.Status)
		} else {
			fmt.Printf(Green+"[+] %s found\n"+Reset, path)
			fmt.Println("-> Found through a direct request")
			fmt.Println("")
		}
	}
}

func testRobotxt(url string) {
	robotURL := url + "/robots.txt"
	resp, err := http.Get(robotURL)
	if err != nil {
		fmt.Println(Red + "[-] robots.txt not accessible" + Reset)
		fmt.Println("")
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound || resp.StatusCode == http.StatusForbidden {
		fmt.Println(Red + "[-] robots.txt not accessible" + Reset)
		fmt.Println("")
	} else {
		fmt.Println(Green + "[+] robots.txt found" + Reset)
		fmt.Println("-> Found through a direct request")
		fmt.Println("-> Location: " + url + "robots.txt")
		fmt.Println("")
	}
}

func testXMLRPC(url string) {
	xmlrpcURL := url + "/xmlrpc.php"
	resp, err := http.Get(xmlrpcURL)
	if err != nil {
		fmt.Println(Red+"[-] XMLRPC not accessible:"+Reset, err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound || resp.StatusCode == http.StatusForbidden {
		fmt.Println(Red+"[-] XMLRPC not accessible:"+Reset, resp.Status)
		fmt.Println("")
	} else {
		fmt.Println(Green + "[+] XMLRPC found" + Reset)
		fmt.Println("-> Found through a direct request")
		fmt.Println("-> Location: " + url + "xmlrpc.php")
		fmt.Println("-> Reference: https://codex.wordpress.org/XML-RPC_WordPress_API")
		fmt.Println(Cyan + "-> Possible exploit: https://nitesculucian.github.io/2019/07/01/exploiting-the-xmlrpc-php-on-all-wordpress-versions/" + Reset)
		fmt.Println("")
	}
}

func testThemeLoc(url string) {
	themeURL := url + "/wp-content/themes/"
	resp, err := http.Get(themeURL)
	if err != nil {
		fmt.Println(Red+"[-] Themes not accessible:"+Reset, err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotFound || resp.StatusCode == http.StatusForbidden {
		fmt.Println(Red + "[-] Themes not accessible:" + Reset)
		fmt.Println("")
	} else {
		fmt.Println(Green + "[+] Themes dir found" + Reset)
		fmt.Println("-> Found through a direct request")
		fmt.Println("-> Location: " + url + "/wp-content/themes/")
		fmt.Println("-> Reference: https://codex.wordpress.org/Theme_Development")
		fmt.Println("")
	}
}
