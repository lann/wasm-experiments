#!/usr/bin/env python2

import sys
import threading
import webbrowser

import BaseHTTPServer, SimpleHTTPServer

PORT = 8080

def open_browser():
    url = 'http://127.0.0.1:%d' % PORT
    print(url)
    webbrowser.open_new(url)
threading.Timer(1, open_browser).start()

Handler = SimpleHTTPServer.SimpleHTTPRequestHandler
Handler.extensions_map['.wasm'] = 'application/wasm'
httpd = BaseHTTPServer.HTTPServer(('127.0.0.1', PORT), Handler)
httpd.serve_forever()
