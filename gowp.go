package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	siteURL := flag.String("u", "", "Website URL")
	flag.Parse()

	if *siteURL == "" {
		fmt.Println("Please specify the website ( https://example.com ) using the -u argument")
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
	fmt.Println("By MALTA! | I am not responsible for your actions, this script was for educational purposes only.")
	fmt.Println("BTC : bc1qa454njt7dd8l505puwudwe2wrvcztxd5d44azg | XMR : 46xP9bvZnoR2AUGd6yrpq9RC3XuLJm5KrGT8GtCUrhCtcBJ5M1nnD1gWAnnD6tLtWqKeoJi8UuNsk41K8kE8UYkdTaosjGh")
	fmt.Println("Session ID : 05f643faf9020221785ecebbef7df21b4c10491e3c07795e264a04132f204c161f")
	fmt.Println("")
	time.Sleep(2 * time.Second)
	fmt.Println("[+] The scan is in progress... Please wait.")
	pingWebsite(*siteURL)
	testWpAdmin(*siteURL)
	testUpload(*siteURL)
	testSitemap(*siteURL)
	testSitemapPost(*siteURL)
	testSitemapPost2(*siteURL)
	testSitemapPage(*siteURL)
	testRobotxt(*siteURL)
	testXMLRPC(*siteURL)
}

func pingWebsite(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("[-] The website is not accessible, scan cancelled.")
		os.Exit(1)
	}

	defer resp.Body.Close()

	fmt.Println("[+] The website is online")
	fmt.Println("-> Find with a ping")
	fmt.Println("")
}

func testWpAdmin(url string) {
	adminURL := url + "/wp-admin"
	resp, err := http.Get(adminURL)
	if err != nil {
		fmt.Println("[-] wp-admin not accessible :", err)
		return
	}

	defer resp.Body.Close()

	fmt.Println("[+] wp-admin found")
	fmt.Println("-> Find with a direct request")
	fmt.Println("")
}

func testUpload(url string) {
	uploadURL := url + "/wp-content/uploads/"
	resp, err := http.Get(uploadURL)
	if err != nil {
		fmt.Println("[-] Upload dir not available :", err)
		return
	}

	defer resp.Body.Close()

	fmt.Println("[+] Upload dir found")
	fmt.Println("-> Find with a direct request")
	fmt.Println("")
}

func testSitemap(url string) {
	sitemapURL := url + "/sitemap_index.xml"
	resp, err := http.Get(sitemapURL)
	if err != nil {
		fmt.Println("[-] sitemap_index.xml not accessible :", err)
		return
	}

	defer resp.Body.Close()

	fmt.Println("[+] sitemap_index.xml found")
	fmt.Println("-> Find with a direct request")
	fmt.Println("")
}

func testSitemapPost(url string) {
	sitemapURL := url + "/wp-sitemap-posts-post-1.xml"
	resp, err := http.Get(sitemapURL)
	if err != nil {
		fmt.Println("[-] wp-sitemap-posts-post-1.xml not accessible :", err)
		return
	}

	defer resp.Body.Close()

	fmt.Println("[+] wp-sitemap-posts-post-1.xml found")
	fmt.Println("-> Find with a direct request")
	fmt.Println("")
}

func testSitemapPost2(url string) {
	sitemapURL := url + "/wp-sitemap-posts-post-2.xml"
	resp, err := http.Get(sitemapURL)
	if err != nil {
		fmt.Println("[-] wp-sitemap-posts-post-2.xml not accessible :", err)
		return
	}

	defer resp.Body.Close()

	fmt.Println("[+] wp-sitemap-posts-post-2.xml found")
	fmt.Println("-> Find with a direct request")
	fmt.Println("")
}

func testSitemapPage(url string) {
	sitemapURL := url + "/wp-sitemap-posts-page-1.xml"
	resp, err := http.Get(sitemapURL)
	if err != nil {
		fmt.Println("[-] wp-sitemap-posts-page-1.xml not accessible :", err)
		return
	}

	defer resp.Body.Close()

	fmt.Println("[+] wp-sitemap-posts-page-1.xml found")
	fmt.Println("-> Find with a direct request")
	fmt.Println("")
}

func testRobotxt(url string) {
	sitemapURL := url + "/robot.txt"
	resp, err := http.Get(sitemapURL)
	if err != nil {
		fmt.Println("[-] Robot.txt not accessible :", err)
		return
	}

	defer resp.Body.Close()

	fmt.Println("|+| Robot.txt found")
	fmt.Println("-> Find with a direct request")
	fmt.Println("")
}

func testXMLRPC(url string) {
	sitemapURL := url + "/xmlrpc.php"
	resp, err := http.Get(sitemapURL)
	if err != nil {
		fmt.Println("[-] XMLRPC not accessible :", err)
		return
	}

	defer resp.Body.Close()

	fmt.Println("[+] XMLRPC found")
	fmt.Println("-> Find with a direct request")
	fmt.Println("-> https://codex.wordpress.org/XML-RPC_WordPress_API")
	fmt.Println("-> Possible exploit : https://nitesculucian.github.io/2019/07/01/exploiting-the-xmlrpc-php-on-all-wordpress-versions/")
	fmt.Println("")

}

