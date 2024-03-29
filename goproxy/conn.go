package goproxy

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"net/textproto"
	"net/url"
	"strings"
)

type conn struct {
	rwc    net.Conn
	brc    *bufio.Reader
	server *Server
}

func (c *conn) serve() {
	defer c.rwc.Close()
	rawHttpRequestHeader, remote, isHttps, err := c.getTunnelInfo()
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("connecting to " + remote)

	remoteConn, err := net.Dial("tcp", remote)
	if err != nil {
		log.Println(err)
		return
	}

	if isHttps {
		log.Println("get https req remote %s \n", remote)
		_, err = c.rwc.Write([]byte("HTTP/1.1 200 Connection established\r\n\r\n"))
		if err != nil {
			log.Println(err)
			return
		}
	} else {
		_, err = rawHttpRequestHeader.WriteTo(remoteConn)
		if err != nil {
			log.Println(err)
			return
		}
	}

	log.Println("begin", c.rwc.RemoteAddr(), "<->", remote)
	c.tunnel(remoteConn)
	log.Println("stop", c.rwc.RemoteAddr(), "<->", remote)

}

func (c *conn) getTunnelInfo() (rawReqHeader bytes.Buffer, host string, isHttps bool, err error) {
	tp := textproto.NewReader(c.brc)

	// First line: GET /index.html HTTP/1.0
	var requestLine string
	if requestLine, err = tp.ReadLine(); err != nil {
		return
	}

	method, requestURI, _, ok := parseRequestLine(requestLine)
	if !ok {
		err = &BadRequestError{"malformed HTTP request"}
		return
	}

	// https request
	if method == "CONNECT" {
		isHttps = true
		requestURI = "http://" + requestURI
	}

	// get remote host
	uriInfo, err := url.ParseRequestURI(requestURI)
	if err != nil {
		return
	}

	// lines: Key: value.
	mimeHeader, err := tp.ReadMIMEHeader()
	if err != nil {
		return
	}

	if uriInfo.Host == "" {
		host = mimeHeader.Get("Host")
	} else {
		if strings.Index(uriInfo.Host, ":") == -1 {
			host = uriInfo.Host + ":80"
		} else {
			host = uriInfo.Host
		}
	}

	rawReqHeader.WriteString(requestLine + "\r\n")
	for k, vs := range mimeHeader {
		for _, v := range vs {
			rawReqHeader.WriteString(fmt.Sprintf("%s: %s\r\n", k, v))
		}
	}
	rawReqHeader.WriteString("\r\n")
	return
}

func (c *conn) tunnel(remoteConn net.Conn) {
	go func() {
		_, err := c.brc.WriteTo(remoteConn)
		if err != nil {
			log.Println(err)
		}
		remoteConn.Close()
	}()
	_, err := io.Copy(c.rwc, remoteConn)
	if err != nil {
		log.Println(err)
	}
}

func parseRequestLine(line string) (method, requestURI, proto string, ok bool) {
	s1 := strings.Index(line, " ")
	s2 := strings.Index(line[s1+1:], " ")
	if s1 < 0 || s2 < 0 {
		return
	}
	s2 += s1 + 1
	return line[:s1], line[s1+1 : s2], line[s2+1:], true
}

type BadRequestError struct {
	what string
}

func (b *BadRequestError) Error() string {
	return b.what
}
