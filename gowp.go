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
	fmt.Println("By MALTA! | I am not responsible for your actions, this script was for educational purposes only.")
	fmt.Println("BTC: bc1qa454njt7dd8l505puwudwe2wrvcztxd5d44azg | XMR: 46xP9bvZnoR2AUGd6yrpq9RC3XuLJm5KrGT8GtCUrhCtcBJ5M1nnD1gWAnnD6tLtWqKeoJi8UuNsk41K8kE8UYkdTaosjGh")
	fmt.Println("Session ID: 05f643faf9020221785ecebbef7df21b4c10491e3c07795e264a04132f204c161f")
	fmt.Println("")
	time.Sleep(2 * time.Second)
	fmt.Println("[INFO] The scan is in progress... Please wait.")
	fmt.Println("")
	pingWebsite(*siteURL)
	testWpAdmin(*siteURL)
	testUpload(*siteURL)
	fmt.Println("[INFO] Scan for sitemaps in progress...")
	fmt.Println("")
	testSitemap(*siteURL)
	fmt.Println("[INFO] Scan for interesting files in progress...")
	fmt.Println("")
	testRobotxt(*siteURL)
	testXMLRPC(*siteURL)
}

func pingWebsite(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("[-] The website is not accessible, scan canceled.")
		os.Exit(1)
	}

	defer resp.Body.Close()

	fmt.Println("[+] The website is online")
	fmt.Println("-> Found with a ping")
	fmt.Println("")
}

func testWpAdmin(url string) {
	adminURL := url + "/wp-admin"
	resp, err := http.Get(adminURL)
	if err != nil {
		fmt.Println("Error accessing wp-admin")
		fmt.Println("")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound || resp.StatusCode == http.StatusForbidden {
		fmt.Println("[-] wp-admin not accessible")
		fmt.Println("")
	} else {
		fmt.Println("[+] wp-admin accessible")
		fmt.Println("-> Location: " + url + "wp-admin")
	}
}

func testUpload(url string) {
	uploadURL := url + "/wp-content/uploads/"
	resp, err := http.Get(uploadURL)
	if err != nil {
		fmt.Println("[-] Upload directory not available")
		fmt.Println("")
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound || resp.StatusCode == http.StatusForbidden {
		fmt.Println("[-] Upload directory not accessible")
		fmt.Println("")
	} else {
		fmt.Println("[+] Upload directory found")
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
			fmt.Printf("[-] %s not accessible: %s\n\n", path, err)
			continue
		}

		defer resp.Body.Close()

		if resp.StatusCode == http.StatusNotFound || resp.StatusCode == http.StatusForbidden {
			fmt.Printf("[-] %s not accessible: %s\n\n", path, resp.Status)
		} else {
			fmt.Printf("[+] %s found\n", path)
			fmt.Println("-> Found through a direct request")
			fmt.Println("")
		}
	}
}

func testRobotxt(url string) {
	robotURL := url + "/robots.txt"
	resp, err := http.Get(robotURL)
	if err != nil {
		fmt.Println("[-] robots.txt not accessible")
		fmt.Println("")
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound || resp.StatusCode == http.StatusForbidden {
		fmt.Println("[-] robots.txt not accessible")
		fmt.Println("")
	} else {
		fmt.Println("[+] robots.txt found")
		fmt.Println("-> Found through a direct request")
		fmt.Println("-> Location: " + url + "robots.txt")
		fmt.Println("")
	}
}

func testXMLRPC(url string) {
	xmlrpcURL := url + "/xmlrpc.php"
	resp, err := http.Get(xmlrpcURL)
	if err != nil {
		fmt.Println("[-] XMLRPC not accessible:", err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound || resp.StatusCode == http.StatusForbidden {
		fmt.Println("[-] XMLRPC not accessible:", resp.Status)
		fmt.Println("")
	} else {
		fmt.Println("[+] XMLRPC found")
		fmt.Println("-> Found through a direct request")
		fmt.Println("-> Location: " + url + "xmlrpc.php")
		fmt.Println("-> Reference: https://codex.wordpress.org/XML-RPC_WordPress_API")
		fmt.Println("-> Possible exploit: https://nitesculucian.github.io/2019/07/01/exploiting-the-xmlrpc-php-on-all-wordpress-versions/")
		fmt.Println("")
	}
}
