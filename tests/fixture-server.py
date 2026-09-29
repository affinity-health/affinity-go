"""Loopback-only synthetic fixture server for the SDK contract tests."""
import json
import os
from http.server import BaseHTTPRequestHandler, ThreadingHTTPServer
from pathlib import Path
from urllib.parse import urlsplit, unquote, parse_qsl
import re

fixtures = json.loads(Path(__file__).with_name('fixture.json').read_text())
traces = {}
failures = set()
class Handler(BaseHTTPRequestHandler):
    def log_message(self, *args):
        pass
    def respond(self, data, status=200):
        payload = json.dumps(data).encode()
        self.send_response(status)
        self.send_header('Content-Type', 'application/json')
        self.send_header('Content-Length', str(len(payload)))
        if status == 429: self.send_header('Retry-After', '0')
        self.end_headers()
        self.wfile.write(payload)
    def do_GET(self): self.handle_request()
    def do_POST(self): self.handle_request()
    def do_PATCH(self): self.handle_request()
    def do_PUT(self): self.handle_request()
    def do_DELETE(self): self.handle_request()
    def handle_request(self):
        url = urlsplit(self.path)
        if url.path == '/healthz': return self.respond({'ok': True})
        parts = url.path.split('/')
        case = parts[1]
        path = unquote('/' + '/'.join(parts[2:]))
        if path == '/trace': return self.respond(traces.get(case, []))
        if path == '/reset':
            traces.pop(case, None)
            failures.discard(case)
            return self.respond({'ok': True})
        raw = self.rfile.read(int(self.headers.get('Content-Length', '0')))
        body = json.loads(raw) if raw else None
        query = dict(parse_qsl(url.query))
        traces.setdefault(case, []).append({'path': path, 'method': self.command, 'query': query, 'body': body, 'key': self.headers.get('Idempotency-Key'), 'version': self.headers.get('Affinity-Version')})
        if path == '/v1/auth/access':
            practice = 'platform' not in case
            return self.respond({'object': 'api_access', 'apiKey': {'id': 'key_a', 'object': 'api_key', 'keyPrefix': 'test'}, 'livemode': False, 'scopes': ['patients:read'], 'serviceAccount': {'id': 'sa_a', 'object': 'service_account', 'apiVersion': '2026-09-28', 'subjectType': 'practice' if practice else 'platform', 'subjectId': 'prac_a' if practice else 'acct_a'}})
        if 'pat_error' in path or ('retry' in case and self.command == 'PATCH' and case not in failures):
            failures.add(case)
            return self.respond({'type': 'about:blank', 'title': 'Rate limited', 'status': 429, 'detail': 'private detail', 'code': 'rate_limited', 'requestId': 'req_a', 'instance': path}, 429)
        for op in fixtures:
            if op['verb'] != self.command or not re.fullmatch(re.sub(r'\{[^}]+\}', '[^/]+', op['path']), path): continue
            result = json.loads(json.dumps(op['response']))
            if isinstance(result, dict):
                if 'id' in result: result['id'] = 'ord_a' if op['group'].startswith('orders') else 'pat_a'
                if 'practiceId' in result: result['practiceId'] = 'prac_a'
                if op['paginated']:
                    item = json.loads(json.dumps(op['item']))
                    item['id'] = 'pat_b' if query.get('startingAfter') else 'pat_a'
                    result.update(data=[item], hasMore=not bool(query.get('startingAfter')))
            return self.respond(result)
        self.respond({'error': 'Unknown fixture route'}, 404)

if __name__ == '__main__':
    ThreadingHTTPServer(('127.0.0.1', int(os.environ.get('SDK_TEST_PORT', '5199'))), Handler).serve_forever()
