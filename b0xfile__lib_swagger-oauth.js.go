// Code generaTed by fileb0x at "2018-07-11 17:48:01.885649322 +0800 CST m=+0.592184511" from config file "b0x.yml" DO NOT EDIT.
// modified(2018-07-11 14:52:04 +0800 CST)
// original path: dist/lib/swagger-oauth.js

package swaggerui

import (
	"os"
)

// FileLibSwaggerOauthJs is "/lib/swagger-oauth.js"
var FileLibSwaggerOauthJs = []byte("\x76\x61\x72\x20\x61\x70\x70\x4e\x61\x6d\x65\x3b\x0a\x76\x61\x72\x20\x70\x6f\x70\x75\x70\x4d\x61\x73\x6b\x3b\x0a\x76\x61\x72\x20\x70\x6f\x70\x75\x70\x44\x69\x61\x6c\x6f\x67\x3b\x0a\x76\x61\x72\x20\x63\x6c\x69\x65\x6e\x74\x49\x64\x3b\x0a\x76\x61\x72\x20\x72\x65\x61\x6c\x6d\x3b\x0a\x76\x61\x72\x20\x6f\x61\x75\x74\x68\x32\x4b\x65\x79\x4e\x61\x6d\x65\x3b\x0a\x76\x61\x72\x20\x72\x65\x64\x69\x72\x65\x63\x74\x5f\x75\x72\x69\x3b\x0a\x0a\x66\x75\x6e\x63\x74\x69\x6f\x6e\x20\x68\x61\x6e\x64\x6c\x65\x4c\x6f\x67\x69\x6e\x28\x29\x20\x7b\x0a\x20\x20\x76\x61\x72\x20\x73\x63\x6f\x70\x65\x73\x20\x3d\x20\x5b\x5d\x3b\x0a\x0a\x20\x20\x76\x61\x72\x20\x61\x75\x74\x68\x73\x20\x3d\x20\x77\x69\x6e\x64\x6f\x77\x2e\x73\x77\x61\x67\x67\x65\x72\x55\x69\x2e\x61\x70\x69\x2e\x61\x75\x74\x68\x53\x63\x68\x65\x6d\x65\x73\x20\x7c\x7c\x20\x77\x69\x6e\x64\x6f\x77\x2e\x73\x77\x61\x67\x67\x65\x72\x55\x69\x2e\x61\x70\x69\x2e\x73\x65\x63\x75\x72\x69\x74\x79\x44\x65\x66\x69\x6e\x69\x74\x69\x6f\x6e\x73\x3b\x0a\x20\x20\x69\x66\x20\x28\x61\x75\x74\x68\x73\x29\x20\x7b\x0a\x20\x20\x20\x20\x76\x61\x72\x20\x6b\x65\x79\x3b\x0a\x20\x20\x20\x20\x76\x61\x72\x20\x64\x65\x66\x73\x20\x3d\x20\x61\x75\x74\x68\x73\x3b\x0a\x20\x20\x20\x20\x66\x6f\x72\x20\x28\x6b\x65\x79\x20\x69\x6e\x20\x64\x65\x66\x73\x29\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x76\x61\x72\x20\x61\x75\x74\x68\x20\x3d\x20\x64\x65\x66\x73\x5b\x6b\x65\x79\x5d\x3b\x0a\x20\x20\x20\x20\x20\x20\x69\x66\x20\x28\x61\x75\x74\x68\x2e\x74\x79\x70\x65\x20\x3d\x3d\x3d\x20\x27\x6f\x61\x75\x74\x68\x32\x27\x20\x26\x26\x20\x61\x75\x74\x68\x2e\x73\x63\x6f\x70\x65\x73\x29\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x6f\x61\x75\x74\x68\x32\x4b\x65\x79\x4e\x61\x6d\x65\x20\x3d\x20\x6b\x65\x79\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x76\x61\x72\x20\x73\x63\x6f\x70\x65\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x69\x66\x20\x28\x41\x72\x72\x61\x79\x2e\x69\x73\x41\x72\x72\x61\x79\x28\x61\x75\x74\x68\x2e\x73\x63\x6f\x70\x65\x73\x29\x29\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x2f\x2f\x20\x31\x2e\x32\x20\x73\x75\x70\x70\x6f\x72\x74\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x76\x61\x72\x20\x69\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x66\x6f\x72\x20\x28\x69\x20\x3d\x20\x30\x3b\x20\x69\x20\x3c\x20\x61\x75\x74\x68\x2e\x73\x63\x6f\x70\x65\x73\x2e\x6c\x65\x6e\x67\x74\x68\x3b\x20\x69\x2b\x2b\x29\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x73\x63\x6f\x70\x65\x73\x2e\x70\x75\x73\x68\x28\x61\x75\x74\x68\x2e\x73\x63\x6f\x70\x65\x73\x5b\x69\x5d\x29\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x7d\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x7d\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x65\x6c\x73\x65\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x2f\x2f\x20\x32\x2e\x30\x20\x73\x75\x70\x70\x6f\x72\x74\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x66\x6f\x72\x20\x28\x73\x63\x6f\x70\x65\x20\x69\x6e\x20\x61\x75\x74\x68\x2e\x73\x63\x6f\x70\x65\x73\x29\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x73\x63\x6f\x70\x65\x73\x2e\x70\x75\x73\x68\x28\x7b\x73\x63\x6f\x70\x65\x3a\x20\x73\x63\x6f\x70\x65\x2c\x20\x64\x65\x73\x63\x72\x69\x70\x74\x69\x6f\x6e\x3a\x20\x61\x75\x74\x68\x2e\x73\x63\x6f\x70\x65\x73\x5b\x73\x63\x6f\x70\x65\x5d\x7d\x29\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x7d\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x7d\x0a\x20\x20\x20\x20\x20\x20\x7d\x0a\x20\x20\x20\x20\x7d\x0a\x20\x20\x7d\x0a\x0a\x20\x20\x69\x66\x20\x28\x77\x69\x6e\x64\x6f\x77\x2e\x73\x77\x61\x67\x67\x65\x72\x55\x69\x2e\x61\x70\x69\x0a\x20\x20\x20\x20\x26\x26\x20\x77\x69\x6e\x64\x6f\x77\x2e\x73\x77\x61\x67\x67\x65\x72\x55\x69\x2e\x61\x70\x69\x2e\x69\x6e\x66\x6f\x29\x20\x7b\x0a\x20\x20\x20\x20\x61\x70\x70\x4e\x61\x6d\x65\x20\x3d\x20\x77\x69\x6e\x64\x6f\x77\x2e\x73\x77\x61\x67\x67\x65\x72\x55\x69\x2e\x61\x70\x69\x2e\x69\x6e\x66\x6f\x2e\x74\x69\x74\x6c\x65\x3b\x0a\x20\x20\x7d\x0a\x0a\x20\x20\x70\x6f\x70\x75\x70\x44\x69\x61\x6c\x6f\x67\x20\x3d\x20\x24\x28\x0a\x20\x20\x20\x20\x5b\x0a\x20\x20\x20\x20\x20\x20\x27\x3c\x64\x69\x76\x20\x63\x6c\x61\x73\x73\x3d\x22\x6d\x6f\x64\x61\x6c\x20\x61\x70\x69\x2d\x70\x6f\x70\x75\x70\x2d\x64\x69\x61\x6c\x6f\x67\x20\x69\x6e\x22\x20\x69\x64\x3d\x22\x63\x72\x65\x64\x65\x6e\x74\x69\x61\x6c\x73\x2d\x6d\x6f\x64\x61\x6c\x22\x20\x74\x61\x62\x69\x6e\x64\x65\x78\x3d\x22\x2d\x31\x22\x20\x72\x6f\x6c\x65\x3d\x22\x64\x69\x61\x6c\x6f\x67\x22\x20\x61\x72\x69\x61\x2d\x6c\x61\x62\x65\x6c\x6c\x65\x64\x62\x79\x3d\x22\x63\x72\x65\x64\x65\x6e\x74\x69\x61\x6c\x73\x2d\x6d\x6f\x64\x61\x6c\x2d\x6c\x61\x62\x65\x6c\x22\x20\x61\x72\x69\x61\x2d\x68\x69\x64\x64\x65\x6e\x3d\x22\x66\x61\x6c\x73\x65\x22\x20\x73\x74\x79\x6c\x65\x3d\x22\x64\x69\x73\x70\x6c\x61\x79\x3a\x20\x62\x6c\x6f\x63\x6b\x3b\x22\x3e\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x27\x3c\x64\x69\x76\x20\x63\x6c\x61\x73\x73\x3d\x22\x6d\x6f\x64\x61\x6c\x2d\x64\x69\x61\x6c\x6f\x67\x22\x3e\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x27\x3c\x64\x69\x76\x20\x63\x6c\x61\x73\x73\x3d\x22\x6d\x6f\x64\x61\x6c\x2d\x63\x6f\x6e\x74\x65\x6e\x74\x22\x3e\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x27\x3c\x64\x69\x76\x20\x63\x6c\x61\x73\x73\x3d\x22\x6d\x6f\x64\x61\x6c\x2d\x68\x65\x61\x64\x65\x72\x22\x3e\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x27\x3c\x62\x75\x74\x74\x6f\x6e\x20\x74\x79\x70\x65\x3d\x22\x62\x75\x74\x74\x6f\x6e\x22\x20\x63\x6c\x61\x73\x73\x3d\x22\x63\x6c\x6f\x73\x65\x20\x61\x70\x69\x2d\x70\x6f\x70\x75\x70\x2d\x63\x61\x6e\x63\x65\x6c\x22\x20\x64\x61\x74\x61\x2d\x64\x69\x73\x6d\x69\x73\x73\x3d\x22\x6d\x6f\x64\x61\x6c\x22\x3e\x3c\x73\x70\x61\x6e\x20\x61\x72\x69\x61\x2d\x68\x69\x64\x64\x65\x6e\x3d\x22\x74\x72\x75\x65\x22\x3e\xc3\x97\x3c\x2f\x73\x70\x61\x6e\x3e\x3c\x73\x70\x61\x6e\x20\x63\x6c\x61\x73\x73\x3d\x22\x73\x72\x2d\x6f\x6e\x6c\x79\x22\x3e\x43\x6c\x6f\x73\x65\x3c\x2f\x73\x70\x61\x6e\x3e\x3c\x2f\x62\x75\x74\x74\x6f\x6e\x3e\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x27\x3c\x68\x33\x20\x63\x6c\x61\x73\x73\x3d\x22\x6d\x6f\x64\x61\x6c\x2d\x74\x69\x74\x6c\x65\x22\x20\x69\x64\x3d\x22\x63\x72\x65\x64\x65\x6e\x74\x69\x61\x6c\x73\x2d\x6d\x6f\x64\x61\x6c\x2d\x6c\x61\x62\x65\x6c\x22\x3e\x53\x65\x6c\x65\x63\x74\x20\x4f\x41\x75\x74\x68\x32\x2e\x30\x20\x53\x63\x6f\x70\x65\x73\x3c\x2f\x68\x33\x3e\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x27\x3c\x2f\x64\x69\x76\x3e\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x27\x3c\x64\x69\x76\x20\x63\x6c\x61\x73\x73\x3d\x22\x6d\x6f\x64\x61\x6c\x2d\x62\x6f\x64\x79\x22\x3e\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x27\x3c\x70\x3e\x53\x63\x6f\x70\x65\x73\x20\x61\x72\x65\x20\x75\x73\x65\x64\x20\x74\x6f\x20\x67\x72\x61\x6e\x74\x20\x61\x6e\x20\x61\x70\x70\x6c\x69\x63\x61\x74\x69\x6f\x6e\x20\x64\x69\x66\x66\x65\x72\x65\x6e\x74\x20\x6c\x65\x76\x65\x6c\x73\x20\x6f\x66\x20\x61\x63\x63\x65\x73\x73\x20\x74\x6f\x20\x64\x61\x74\x61\x20\x6f\x6e\x20\x62\x65\x68\x61\x6c\x66\x20\x6f\x66\x20\x74\x68\x65\x20\x65\x6e\x64\x20\x75\x73\x65\x72\x2e\x20\x45\x61\x63\x68\x20\x41\x50\x49\x20\x6d\x61\x79\x20\x64\x65\x63\x6c\x61\x72\x65\x20\x6f\x6e\x65\x20\x6f\x72\x20\x6d\x6f\x72\x65\x20\x73\x63\x6f\x70\x65\x73\x2e\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x27\x3c\x61\x20\x68\x72\x65\x66\x3d\x22\x23\x22\x3e\x4c\x65\x61\x72\x6e\x20\x68\x6f\x77\x20\x74\x6f\x20\x75\x73\x65\x3c\x2f\x61\x3e\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x27\x3c\x2f\x70\x3e\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x27\x3c\x70\x3e\x3c\x73\x74\x72\x6f\x6e\x67\x3e\x27\x20\x2b\x20\x61\x70\x70\x4e\x61\x6d\x65\x20\x2b\x20\x27\x3c\x2f\x73\x74\x72\x6f\x6e\x67\x3e\x20\x41\x50\x49\x20\x72\x65\x71\x75\x69\x72\x65\x73\x20\x74\x68\x65\x20\x66\x6f\x6c\x6c\x6f\x77\x69\x6e\x67\x20\x73\x63\x6f\x70\x65\x73\x2e\x20\x53\x65\x6c\x65\x63\x74\x20\x77\x68\x69\x63\x68\x20\x6f\x6e\x65\x73\x20\x79\x6f\x75\x20\x77\x61\x6e\x74\x20\x74\x6f\x20\x67\x72\x61\x6e\x74\x20\x74\x6f\x20\x53\x77\x61\x67\x67\x65\x72\x20\x55\x49\x2e\x3c\x2f\x70\x3e\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x27\x3c\x66\x6f\x72\x6d\x3e\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x27\x3c\x64\x69\x76\x20\x63\x6c\x61\x73\x73\x3d\x22\x61\x70\x69\x2d\x70\x6f\x70\x75\x70\x2d\x73\x63\x6f\x70\x65\x73\x22\x3e\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x27\x3c\x64\x69\x76\x20\x63\x6c\x61\x73\x73\x3d\x22\x73\x63\x6f\x70\x65\x73\x22\x3e\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x27\x3c\x2f\x64\x69\x76\x3e\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x27\x3c\x2f\x64\x69\x76\x3e\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x27\x3c\x66\x6f\x72\x6d\x3e\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x27\x3c\x70\x20\x63\x6c\x61\x73\x73\x3d\x22\x65\x72\x72\x6f\x72\x2d\x6d\x73\x67\x22\x3e\x3c\x2f\x70\x3e\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x27\x3c\x2f\x64\x69\x76\x3e\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x27\x3c\x64\x69\x76\x20\x63\x6c\x61\x73\x73\x3d\x22\x6d\x6f\x64\x61\x6c\x2d\x66\x6f\x6f\x74\x65\x72\x22\x3e\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x27\x3c\x64\x69\x76\x20\x63\x6c\x61\x73\x73\x3d\x22\x61\x70\x69\x2d\x70\x6f\x70\x75\x70\x2d\x61\x63\x74\x69\x6f\x6e\x73\x22\x3e\x3c\x62\x75\x74\x74\x6f\x6e\x20\x63\x6c\x61\x73\x73\x3d\x22\x61\x70\x69\x2d\x70\x6f\x70\x75\x70\x2d\x63\x61\x6e\x63\x65\x6c\x20\x62\x74\x6e\x20\x62\x74\x6e\x2d\x64\x65\x66\x61\x75\x6c\x74\x22\x20\x74\x79\x70\x65\x3d\x22\x62\x75\x74\x74\x6f\x6e\x22\x3e\x43\x61\x6e\x63\x65\x6c\x3c\x2f\x62\x75\x74\x74\x6f\x6e\x3e\x3c\x62\x75\x74\x74\x6f\x6e\x20\x63\x6c\x61\x73\x73\x3d\x22\x61\x70\x69\x2d\x70\x6f\x70\x75\x70\x2d\x61\x75\x74\x68\x62\x74\x6e\x20\x62\x74\x6e\x20\x62\x74\x6e\x2d\x70\x72\x69\x6d\x61\x72\x79\x22\x20\x74\x79\x70\x65\x3d\x22\x62\x75\x74\x74\x6f\x6e\x22\x3e\x41\x75\x74\x68\x6f\x72\x69\x7a\x65\x3c\x2f\x62\x75\x74\x74\x6f\x6e\x3e\x3c\x2f\x64\x69\x76\x3e\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x27\x3c\x2f\x64\x69\x76\x3e\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x27\x3c\x2f\x64\x69\x76\x3e\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x27\x3c\x2f\x64\x69\x76\x3e\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x27\x3c\x2f\x64\x69\x76\x3e\x27\x5d\x2e\x6a\x6f\x69\x6e\x28\x27\x27\x29\x29\x3b\x0a\x20\x20\x24\x28\x64\x6f\x63\x75\x6d\x65\x6e\x74\x2e\x62\x6f\x64\x79\x29\x2e\x61\x70\x70\x65\x6e\x64\x28\x70\x6f\x70\x75\x70\x44\x69\x61\x6c\x6f\x67\x29\x3b\x0a\x0a\x20\x20\x70\x6f\x70\x75\x70\x20\x3d\x20\x70\x6f\x70\x75\x70\x44\x69\x61\x6c\x6f\x67\x2e\x66\x69\x6e\x64\x28\x27\x2e\x73\x63\x6f\x70\x65\x73\x27\x29\x2e\x65\x6d\x70\x74\x79\x28\x29\x3b\x0a\x20\x20\x66\x6f\x72\x20\x28\x69\x20\x3d\x20\x30\x3b\x20\x69\x20\x3c\x20\x73\x63\x6f\x70\x65\x73\x2e\x6c\x65\x6e\x67\x74\x68\x3b\x20\x69\x2b\x2b\x29\x20\x7b\x0a\x20\x20\x20\x20\x73\x63\x6f\x70\x65\x20\x3d\x20\x73\x63\x6f\x70\x65\x73\x5b\x69\x5d\x3b\x0a\x20\x20\x20\x20\x73\x74\x72\x20\x3d\x20\x27\x3c\x73\x70\x61\x6e\x20\x64\x61\x74\x61\x2d\x74\x6f\x67\x67\x6c\x65\x2d\x73\x63\x6f\x70\x65\x3d\x22\x27\x20\x2b\x20\x73\x63\x6f\x70\x65\x2e\x73\x63\x6f\x70\x65\x20\x2b\x20\x27\x22\x20\x63\x6c\x61\x73\x73\x3d\x22\x73\x63\x6f\x70\x65\x22\x3e\x27\x20\x2b\x20\x73\x63\x6f\x70\x65\x2e\x73\x63\x6f\x70\x65\x20\x2b\x20\x27\x3c\x2f\x73\x70\x61\x6e\x3e\x27\x3b\x0a\x20\x20\x20\x20\x70\x6f\x70\x75\x70\x2e\x61\x70\x70\x65\x6e\x64\x28\x73\x74\x72\x29\x3b\x0a\x20\x20\x7d\x0a\x0a\x20\x20\x70\x6f\x70\x75\x70\x44\x69\x61\x6c\x6f\x67\x2e\x66\x69\x6e\x64\x28\x27\x73\x63\x6f\x70\x65\x73\x27\x29\x2e\x63\x6c\x69\x63\x6b\x28\x66\x75\x6e\x63\x74\x69\x6f\x6e\x20\x28\x29\x20\x7b\x0a\x20\x20\x20\x20\x70\x6f\x70\x75\x70\x4d\x61\x73\x6b\x2e\x68\x69\x64\x65\x28\x29\x3b\x0a\x20\x20\x20\x20\x70\x6f\x70\x75\x70\x44\x69\x61\x6c\x6f\x67\x2e\x68\x69\x64\x65\x28\x29\x3b\x0a\x20\x20\x20\x20\x70\x6f\x70\x75\x70\x44\x69\x61\x6c\x6f\x67\x2e\x65\x6d\x70\x74\x79\x28\x29\x3b\x0a\x20\x20\x20\x20\x70\x6f\x70\x75\x70\x44\x69\x61\x6c\x6f\x67\x20\x3d\x20\x5b\x5d\x3b\x0a\x20\x20\x7d\x29\x3b\x0a\x0a\x20\x20\x70\x6f\x70\x75\x70\x44\x69\x61\x6c\x6f\x67\x2e\x66\x69\x6e\x64\x28\x27\x5b\x64\x61\x74\x61\x2d\x74\x6f\x67\x67\x6c\x65\x2d\x73\x63\x6f\x70\x65\x5d\x27\x29\x2e\x63\x6c\x69\x63\x6b\x28\x66\x75\x6e\x63\x74\x69\x6f\x6e\x20\x28\x29\x20\x7b\x0a\x20\x20\x20\x20\x24\x28\x74\x68\x69\x73\x29\x2e\x68\x61\x73\x43\x6c\x61\x73\x73\x28\x22\x61\x63\x74\x69\x76\x65\x22\x29\x20\x3f\x20\x24\x28\x74\x68\x69\x73\x29\x2e\x72\x65\x6d\x6f\x76\x65\x43\x6c\x61\x73\x73\x28\x27\x61\x63\x74\x69\x76\x65\x27\x29\x20\x3a\x20\x24\x28\x74\x68\x69\x73\x29\x2e\x61\x64\x64\x43\x6c\x61\x73\x73\x28\x27\x61\x63\x74\x69\x76\x65\x27\x29\x3b\x0a\x20\x20\x7d\x29\x3b\x0a\x0a\x20\x20\x70\x6f\x70\x75\x70\x44\x69\x61\x6c\x6f\x67\x2e\x66\x69\x6e\x64\x28\x27\x62\x75\x74\x74\x6f\x6e\x2e\x61\x70\x69\x2d\x70\x6f\x70\x75\x70\x2d\x63\x61\x6e\x63\x65\x6c\x27\x29\x2e\x63\x6c\x69\x63\x6b\x28\x66\x75\x6e\x63\x74\x69\x6f\x6e\x20\x28\x29\x20\x7b\x0a\x20\x20\x20\x20\x70\x6f\x70\x75\x70\x4d\x61\x73\x6b\x2e\x68\x69\x64\x65\x28\x29\x3b\x0a\x20\x20\x20\x20\x70\x6f\x70\x75\x70\x44\x69\x61\x6c\x6f\x67\x2e\x68\x69\x64\x65\x28\x29\x3b\x0a\x20\x20\x20\x20\x70\x6f\x70\x75\x70\x44\x69\x61\x6c\x6f\x67\x2e\x65\x6d\x70\x74\x79\x28\x29\x3b\x0a\x20\x20\x20\x20\x70\x6f\x70\x75\x70\x44\x69\x61\x6c\x6f\x67\x20\x3d\x20\x5b\x5d\x3b\x0a\x20\x20\x7d\x29\x3b\x0a\x0a\x20\x20\x24\x28\x27\x62\x75\x74\x74\x6f\x6e\x2e\x61\x70\x69\x2d\x70\x6f\x70\x75\x70\x2d\x61\x75\x74\x68\x62\x74\x6e\x27\x29\x2e\x75\x6e\x62\x69\x6e\x64\x28\x29\x3b\x0a\x20\x20\x70\x6f\x70\x75\x70\x44\x69\x61\x6c\x6f\x67\x2e\x66\x69\x6e\x64\x28\x27\x62\x75\x74\x74\x6f\x6e\x2e\x61\x70\x69\x2d\x70\x6f\x70\x75\x70\x2d\x61\x75\x74\x68\x62\x74\x6e\x27\x29\x2e\x63\x6c\x69\x63\x6b\x28\x66\x75\x6e\x63\x74\x69\x6f\x6e\x20\x28\x29\x20\x7b\x0a\x20\x20\x20\x20\x70\x6f\x70\x75\x70\x4d\x61\x73\x6b\x2e\x68\x69\x64\x65\x28\x29\x3b\x0a\x20\x20\x20\x20\x70\x6f\x70\x75\x70\x44\x69\x61\x6c\x6f\x67\x2e\x68\x69\x64\x65\x28\x29\x3b\x0a\x0a\x20\x20\x20\x20\x76\x61\x72\x20\x61\x75\x74\x68\x53\x63\x68\x65\x6d\x65\x73\x20\x3d\x20\x77\x69\x6e\x64\x6f\x77\x2e\x73\x77\x61\x67\x67\x65\x72\x55\x69\x2e\x61\x70\x69\x2e\x61\x75\x74\x68\x53\x63\x68\x65\x6d\x65\x73\x3b\x0a\x20\x20\x20\x20\x76\x61\x72\x20\x68\x6f\x73\x74\x20\x3d\x20\x77\x69\x6e\x64\x6f\x77\x2e\x6c\x6f\x63\x61\x74\x69\x6f\x6e\x3b\x0a\x20\x20\x20\x20\x76\x61\x72\x20\x70\x61\x74\x68\x6e\x61\x6d\x65\x20\x3d\x20\x6c\x6f\x63\x61\x74\x69\x6f\x6e\x2e\x70\x61\x74\x68\x6e\x61\x6d\x65\x2e\x73\x75\x62\x73\x74\x72\x69\x6e\x67\x28\x30\x2c\x20\x6c\x6f\x63\x61\x74\x69\x6f\x6e\x2e\x70\x61\x74\x68\x6e\x61\x6d\x65\x2e\x6c\x61\x73\x74\x49\x6e\x64\x65\x78\x4f\x66\x28\x22\x2f\x22\x29\x29\x3b\x0a\x20\x20\x20\x20\x76\x61\x72\x20\x64\x65\x66\x61\x75\x6c\x74\x52\x65\x64\x69\x72\x65\x63\x74\x55\x72\x6c\x20\x3d\x20\x68\x6f\x73\x74\x2e\x70\x72\x6f\x74\x6f\x63\x6f\x6c\x20\x2b\x20\x27\x2f\x2f\x27\x20\x2b\x20\x68\x6f\x73\x74\x2e\x68\x6f\x73\x74\x20\x2b\x20\x70\x61\x74\x68\x6e\x61\x6d\x65\x20\x2b\x20\x27\x2f\x6f\x32\x63\x2e\x68\x74\x6d\x6c\x27\x3b\x0a\x20\x20\x20\x20\x76\x61\x72\x20\x72\x65\x64\x69\x72\x65\x63\x74\x55\x72\x6c\x20\x3d\x20\x77\x69\x6e\x64\x6f\x77\x2e\x6f\x41\x75\x74\x68\x52\x65\x64\x69\x72\x65\x63\x74\x55\x72\x6c\x20\x7c\x7c\x20\x64\x65\x66\x61\x75\x6c\x74\x52\x65\x64\x69\x72\x65\x63\x74\x55\x72\x6c\x3b\x0a\x20\x20\x20\x20\x76\x61\x72\x20\x75\x72\x6c\x20\x3d\x20\x6e\x75\x6c\x6c\x3b\x0a\x0a\x20\x20\x20\x20\x66\x6f\x72\x20\x28\x76\x61\x72\x20\x6b\x65\x79\x20\x69\x6e\x20\x61\x75\x74\x68\x53\x63\x68\x65\x6d\x65\x73\x29\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x69\x66\x20\x28\x61\x75\x74\x68\x53\x63\x68\x65\x6d\x65\x73\x2e\x68\x61\x73\x4f\x77\x6e\x50\x72\x6f\x70\x65\x72\x74\x79\x28\x6b\x65\x79\x29\x29\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x76\x61\x72\x20\x66\x6c\x6f\x77\x20\x3d\x20\x61\x75\x74\x68\x53\x63\x68\x65\x6d\x65\x73\x5b\x6b\x65\x79\x5d\x2e\x66\x6c\x6f\x77\x3b\x0a\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x69\x66\x20\x28\x61\x75\x74\x68\x53\x63\x68\x65\x6d\x65\x73\x5b\x6b\x65\x79\x5d\x2e\x74\x79\x70\x65\x20\x3d\x3d\x3d\x20\x27\x6f\x61\x75\x74\x68\x32\x27\x20\x26\x26\x20\x66\x6c\x6f\x77\x20\x26\x26\x20\x28\x66\x6c\x6f\x77\x20\x3d\x3d\x3d\x20\x27\x69\x6d\x70\x6c\x69\x63\x69\x74\x27\x20\x7c\x7c\x20\x66\x6c\x6f\x77\x20\x3d\x3d\x3d\x20\x27\x61\x63\x63\x65\x73\x73\x43\x6f\x64\x65\x27\x29\x29\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x76\x61\x72\x20\x64\x65\x74\x73\x20\x3d\x20\x61\x75\x74\x68\x53\x63\x68\x65\x6d\x65\x73\x5b\x6b\x65\x79\x5d\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x75\x72\x6c\x20\x3d\x20\x64\x65\x74\x73\x2e\x61\x75\x74\x68\x6f\x72\x69\x7a\x61\x74\x69\x6f\x6e\x55\x72\x6c\x20\x2b\x20\x27\x3f\x72\x65\x73\x70\x6f\x6e\x73\x65\x5f\x74\x79\x70\x65\x3d\x27\x20\x2b\x20\x28\x66\x6c\x6f\x77\x20\x3d\x3d\x3d\x20\x27\x69\x6d\x70\x6c\x69\x63\x69\x74\x27\x20\x3f\x20\x27\x74\x6f\x6b\x65\x6e\x27\x20\x3a\x20\x27\x63\x6f\x64\x65\x27\x29\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x77\x69\x6e\x64\x6f\x77\x2e\x73\x77\x61\x67\x67\x65\x72\x55\x69\x2e\x74\x6f\x6b\x65\x6e\x4e\x61\x6d\x65\x20\x3d\x20\x64\x65\x74\x73\x2e\x74\x6f\x6b\x65\x6e\x4e\x61\x6d\x65\x20\x7c\x7c\x20\x27\x61\x63\x63\x65\x73\x73\x5f\x74\x6f\x6b\x65\x6e\x27\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x77\x69\x6e\x64\x6f\x77\x2e\x73\x77\x61\x67\x67\x65\x72\x55\x69\x2e\x74\x6f\x6b\x65\x6e\x55\x72\x6c\x20\x3d\x20\x28\x66\x6c\x6f\x77\x20\x3d\x3d\x3d\x20\x27\x61\x63\x63\x65\x73\x73\x43\x6f\x64\x65\x27\x20\x3f\x20\x64\x65\x74\x73\x2e\x74\x6f\x6b\x65\x6e\x55\x72\x6c\x20\x3a\x20\x6e\x75\x6c\x6c\x29\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x7d\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x65\x6c\x73\x65\x20\x69\x66\x20\x28\x61\x75\x74\x68\x53\x63\x68\x65\x6d\x65\x73\x5b\x6b\x65\x79\x5d\x2e\x67\x72\x61\x6e\x74\x54\x79\x70\x65\x73\x29\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x2f\x2f\x20\x31\x2e\x32\x20\x73\x75\x70\x70\x6f\x72\x74\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x76\x61\x72\x20\x6f\x20\x3d\x20\x61\x75\x74\x68\x53\x63\x68\x65\x6d\x65\x73\x5b\x6b\x65\x79\x5d\x2e\x67\x72\x61\x6e\x74\x54\x79\x70\x65\x73\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x66\x6f\x72\x20\x28\x76\x61\x72\x20\x74\x20\x69\x6e\x20\x6f\x29\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x69\x66\x20\x28\x6f\x2e\x68\x61\x73\x4f\x77\x6e\x50\x72\x6f\x70\x65\x72\x74\x79\x28\x74\x29\x20\x26\x26\x20\x74\x20\x3d\x3d\x3d\x20\x27\x69\x6d\x70\x6c\x69\x63\x69\x74\x27\x29\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x76\x61\x72\x20\x64\x65\x74\x73\x20\x3d\x20\x6f\x5b\x74\x5d\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x76\x61\x72\x20\x65\x70\x20\x3d\x20\x64\x65\x74\x73\x2e\x6c\x6f\x67\x69\x6e\x45\x6e\x64\x70\x6f\x69\x6e\x74\x2e\x75\x72\x6c\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x75\x72\x6c\x20\x3d\x20\x64\x65\x74\x73\x2e\x6c\x6f\x67\x69\x6e\x45\x6e\x64\x70\x6f\x69\x6e\x74\x2e\x75\x72\x6c\x20\x2b\x20\x27\x3f\x72\x65\x73\x70\x6f\x6e\x73\x65\x5f\x74\x79\x70\x65\x3d\x74\x6f\x6b\x65\x6e\x27\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x77\x69\x6e\x64\x6f\x77\x2e\x73\x77\x61\x67\x67\x65\x72\x55\x69\x2e\x74\x6f\x6b\x65\x6e\x4e\x61\x6d\x65\x20\x3d\x20\x64\x65\x74\x73\x2e\x74\x6f\x6b\x65\x6e\x4e\x61\x6d\x65\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x7d\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x65\x6c\x73\x65\x20\x69\x66\x20\x28\x6f\x2e\x68\x61\x73\x4f\x77\x6e\x50\x72\x6f\x70\x65\x72\x74\x79\x28\x74\x29\x20\x26\x26\x20\x74\x20\x3d\x3d\x3d\x20\x27\x61\x63\x63\x65\x73\x73\x43\x6f\x64\x65\x27\x29\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x76\x61\x72\x20\x64\x65\x74\x73\x20\x3d\x20\x6f\x5b\x74\x5d\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x76\x61\x72\x20\x65\x70\x20\x3d\x20\x64\x65\x74\x73\x2e\x74\x6f\x6b\x65\x6e\x52\x65\x71\x75\x65\x73\x74\x45\x6e\x64\x70\x6f\x69\x6e\x74\x2e\x75\x72\x6c\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x75\x72\x6c\x20\x3d\x20\x64\x65\x74\x73\x2e\x74\x6f\x6b\x65\x6e\x52\x65\x71\x75\x65\x73\x74\x45\x6e\x64\x70\x6f\x69\x6e\x74\x2e\x75\x72\x6c\x20\x2b\x20\x27\x3f\x72\x65\x73\x70\x6f\x6e\x73\x65\x5f\x74\x79\x70\x65\x3d\x63\x6f\x64\x65\x27\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x77\x69\x6e\x64\x6f\x77\x2e\x73\x77\x61\x67\x67\x65\x72\x55\x69\x2e\x74\x6f\x6b\x65\x6e\x4e\x61\x6d\x65\x20\x3d\x20\x64\x65\x74\x73\x2e\x74\x6f\x6b\x65\x6e\x4e\x61\x6d\x65\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x7d\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x7d\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x7d\x0a\x20\x20\x20\x20\x20\x20\x7d\x0a\x20\x20\x20\x20\x7d\x0a\x20\x20\x20\x20\x76\x61\x72\x20\x73\x63\x6f\x70\x65\x73\x20\x3d\x20\x5b\x5d\x3b\x0a\x20\x20\x20\x20\x76\x61\x72\x20\x6f\x20\x3d\x20\x24\x28\x22\x2e\x73\x63\x6f\x70\x65\x73\x3a\x6c\x61\x73\x74\x20\x2e\x61\x63\x74\x69\x76\x65\x22\x29\x3b\x0a\x20\x20\x20\x20\x66\x6f\x72\x20\x28\x6b\x20\x3d\x20\x30\x3b\x20\x6b\x20\x3c\x20\x6f\x2e\x6c\x65\x6e\x67\x74\x68\x3b\x20\x6b\x2b\x2b\x29\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x76\x61\x72\x20\x73\x63\x6f\x70\x65\x20\x3d\x20\x24\x28\x6f\x5b\x6b\x5d\x29\x2e\x61\x74\x74\x72\x28\x27\x64\x61\x74\x61\x2d\x74\x6f\x67\x67\x6c\x65\x2d\x73\x63\x6f\x70\x65\x27\x29\x3b\x0a\x20\x20\x20\x20\x20\x20\x69\x66\x20\x28\x73\x63\x6f\x70\x65\x73\x2e\x69\x6e\x64\x65\x78\x4f\x66\x28\x73\x63\x6f\x70\x65\x29\x20\x3d\x3d\x3d\x20\x2d\x31\x29\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x73\x63\x6f\x70\x65\x73\x2e\x70\x75\x73\x68\x28\x73\x63\x6f\x70\x65\x29\x3b\x0a\x20\x20\x20\x20\x7d\x0a\x0a\x20\x20\x20\x20\x2f\x2f\x20\x49\x6d\x70\x6c\x69\x63\x69\x74\x20\x61\x75\x74\x68\x20\x72\x65\x63\x6f\x6d\x6d\x65\x6e\x64\x73\x20\x61\x20\x73\x74\x61\x74\x65\x20\x70\x61\x72\x61\x6d\x65\x74\x65\x72\x2e\x0a\x20\x20\x20\x20\x76\x61\x72\x20\x73\x74\x61\x74\x65\x20\x3d\x20\x4d\x61\x74\x68\x2e\x72\x61\x6e\x64\x6f\x6d\x28\x29\x3b\x0a\x0a\x20\x20\x20\x20\x77\x69\x6e\x64\x6f\x77\x2e\x65\x6e\x61\x62\x6c\x65\x64\x53\x63\x6f\x70\x65\x73\x20\x3d\x20\x73\x63\x6f\x70\x65\x73\x3b\x0a\x0a\x20\x20\x20\x20\x72\x65\x64\x69\x72\x65\x63\x74\x5f\x75\x72\x69\x20\x3d\x20\x72\x65\x64\x69\x72\x65\x63\x74\x55\x72\x6c\x3b\x0a\x0a\x20\x20\x20\x20\x75\x72\x6c\x20\x2b\x3d\x20\x27\x26\x72\x65\x64\x69\x72\x65\x63\x74\x5f\x75\x72\x69\x3d\x27\x20\x2b\x20\x65\x6e\x63\x6f\x64\x65\x55\x52\x49\x43\x6f\x6d\x70\x6f\x6e\x65\x6e\x74\x28\x72\x65\x64\x69\x72\x65\x63\x74\x55\x72\x6c\x29\x3b\x0a\x20\x20\x20\x20\x75\x72\x6c\x20\x2b\x3d\x20\x27\x26\x72\x65\x61\x6c\x6d\x3d\x27\x20\x2b\x20\x65\x6e\x63\x6f\x64\x65\x55\x52\x49\x43\x6f\x6d\x70\x6f\x6e\x65\x6e\x74\x28\x72\x65\x61\x6c\x6d\x29\x3b\x0a\x20\x20\x20\x20\x75\x72\x6c\x20\x2b\x3d\x20\x27\x26\x63\x6c\x69\x65\x6e\x74\x5f\x69\x64\x3d\x27\x20\x2b\x20\x65\x6e\x63\x6f\x64\x65\x55\x52\x49\x43\x6f\x6d\x70\x6f\x6e\x65\x6e\x74\x28\x63\x6c\x69\x65\x6e\x74\x49\x64\x29\x3b\x0a\x20\x20\x20\x20\x75\x72\x6c\x20\x2b\x3d\x20\x27\x26\x73\x63\x6f\x70\x65\x3d\x27\x20\x2b\x20\x65\x6e\x63\x6f\x64\x65\x55\x52\x49\x43\x6f\x6d\x70\x6f\x6e\x65\x6e\x74\x28\x73\x63\x6f\x70\x65\x73\x2e\x6a\x6f\x69\x6e\x28\x27\x20\x27\x29\x29\x3b\x0a\x20\x20\x20\x20\x75\x72\x6c\x20\x2b\x3d\x20\x27\x26\x73\x74\x61\x74\x65\x3d\x27\x20\x2b\x20\x65\x6e\x63\x6f\x64\x65\x55\x52\x49\x43\x6f\x6d\x70\x6f\x6e\x65\x6e\x74\x28\x73\x74\x61\x74\x65\x29\x3b\x0a\x0a\x20\x20\x20\x20\x77\x69\x6e\x64\x6f\x77\x2e\x6f\x70\x65\x6e\x28\x75\x72\x6c\x29\x3b\x0a\x20\x20\x7d\x29\x3b\x0a\x0a\x20\x20\x70\x6f\x70\x75\x70\x4d\x61\x73\x6b\x2e\x73\x68\x6f\x77\x28\x29\x3b\x0a\x20\x20\x70\x6f\x70\x75\x70\x44\x69\x61\x6c\x6f\x67\x2e\x73\x68\x6f\x77\x28\x29\x3b\x0a\x7d\x0a\x0a\x66\x75\x6e\x63\x74\x69\x6f\x6e\x20\x68\x61\x6e\x64\x6c\x65\x4c\x6f\x67\x6f\x75\x74\x28\x29\x20\x7b\x0a\x20\x20\x66\x6f\x72\x20\x28\x6b\x65\x79\x20\x69\x6e\x20\x77\x69\x6e\x64\x6f\x77\x2e\x61\x75\x74\x68\x6f\x72\x69\x7a\x61\x74\x69\x6f\x6e\x73\x2e\x61\x75\x74\x68\x7a\x29\x20\x7b\x0a\x20\x20\x20\x20\x77\x69\x6e\x64\x6f\x77\x2e\x61\x75\x74\x68\x6f\x72\x69\x7a\x61\x74\x69\x6f\x6e\x73\x2e\x72\x65\x6d\x6f\x76\x65\x28\x6b\x65\x79\x29\x0a\x20\x20\x7d\x0a\x20\x20\x77\x69\x6e\x64\x6f\x77\x2e\x65\x6e\x61\x62\x6c\x65\x64\x53\x63\x6f\x70\x65\x73\x20\x3d\x20\x6e\x75\x6c\x6c\x3b\x0a\x20\x20\x76\x61\x72\x20\x6f\x61\x75\x74\x68\x42\x74\x6e\x20\x3d\x20\x24\x28\x27\x2e\x61\x70\x69\x2d\x69\x63\x27\x29\x3b\x0a\x20\x20\x6f\x61\x75\x74\x68\x42\x74\x6e\x2e\x61\x64\x64\x43\x6c\x61\x73\x73\x28\x27\x62\x74\x6e\x2d\x64\x65\x66\x61\x75\x6c\x74\x27\x29\x3b\x0a\x20\x20\x6f\x61\x75\x74\x68\x42\x74\x6e\x2e\x72\x65\x6d\x6f\x76\x65\x43\x6c\x61\x73\x73\x28\x27\x62\x74\x6e\x2d\x73\x75\x63\x63\x65\x73\x73\x27\x29\x3b\x0a\x20\x20\x6f\x61\x75\x74\x68\x42\x74\x6e\x2e\x72\x65\x6d\x6f\x76\x65\x43\x6c\x61\x73\x73\x28\x27\x62\x74\x6e\x2d\x77\x61\x72\x6e\x69\x6e\x67\x27\x29\x3b\x0a\x0a\x20\x20\x6f\x61\x75\x74\x68\x42\x74\x6e\x2e\x74\x65\x78\x74\x28\x27\x6f\x61\x75\x74\x68\x27\x29\x3b\x0a\x0a\x20\x20\x24\x28\x27\x23\x69\x6e\x70\x75\x74\x5f\x61\x70\x69\x4b\x65\x79\x27\x29\x2e\x76\x61\x6c\x28\x27\x27\x29\x3b\x0a\x7d\x0a\x0a\x66\x75\x6e\x63\x74\x69\x6f\x6e\x20\x69\x6e\x69\x74\x4f\x41\x75\x74\x68\x28\x6f\x70\x74\x73\x29\x20\x7b\x0a\x20\x20\x76\x61\x72\x20\x6f\x20\x3d\x20\x28\x6f\x70\x74\x73\x20\x7c\x7c\x20\x7b\x7d\x29\x3b\x0a\x20\x20\x76\x61\x72\x20\x65\x72\x72\x6f\x72\x73\x20\x3d\x20\x5b\x5d\x3b\x0a\x0a\x20\x20\x61\x70\x70\x4e\x61\x6d\x65\x20\x3d\x20\x28\x6f\x2e\x61\x70\x70\x4e\x61\x6d\x65\x20\x7c\x7c\x20\x65\x72\x72\x6f\x72\x73\x2e\x70\x75\x73\x68\x28\x27\x6d\x69\x73\x73\x69\x6e\x67\x20\x61\x70\x70\x4e\x61\x6d\x65\x27\x29\x29\x3b\x0a\x20\x20\x70\x6f\x70\x75\x70\x4d\x61\x73\x6b\x20\x3d\x20\x28\x6f\x2e\x70\x6f\x70\x75\x70\x4d\x61\x73\x6b\x20\x7c\x7c\x20\x24\x28\x27\x23\x61\x70\x69\x2d\x63\x6f\x6d\x6d\x6f\x6e\x2d\x6d\x61\x73\x6b\x27\x29\x29\x3b\x0a\x20\x20\x70\x6f\x70\x75\x70\x44\x69\x61\x6c\x6f\x67\x20\x3d\x20\x28\x6f\x2e\x70\x6f\x70\x75\x70\x44\x69\x61\x6c\x6f\x67\x20\x7c\x7c\x20\x24\x28\x27\x2e\x61\x70\x69\x2d\x70\x6f\x70\x75\x70\x2d\x64\x69\x61\x6c\x6f\x67\x27\x29\x29\x3b\x0a\x20\x20\x63\x6c\x69\x65\x6e\x74\x49\x64\x20\x3d\x20\x28\x6f\x2e\x63\x6c\x69\x65\x6e\x74\x49\x64\x20\x7c\x7c\x20\x65\x72\x72\x6f\x72\x73\x2e\x70\x75\x73\x68\x28\x27\x6d\x69\x73\x73\x69\x6e\x67\x20\x63\x6c\x69\x65\x6e\x74\x20\x69\x64\x27\x29\x29\x3b\x0a\x20\x20\x72\x65\x61\x6c\x6d\x20\x3d\x20\x28\x6f\x2e\x72\x65\x61\x6c\x6d\x20\x7c\x7c\x20\x65\x72\x72\x6f\x72\x73\x2e\x70\x75\x73\x68\x28\x27\x6d\x69\x73\x73\x69\x6e\x67\x20\x72\x65\x61\x6c\x6d\x27\x29\x29\x3b\x0a\x0a\x20\x20\x69\x66\x20\x28\x65\x72\x72\x6f\x72\x73\x2e\x6c\x65\x6e\x67\x74\x68\x20\x3e\x20\x30\x29\x20\x7b\x0a\x20\x20\x20\x20\x6c\x6f\x67\x28\x27\x61\x75\x74\x68\x20\x75\x6e\x61\x62\x6c\x65\x20\x69\x6e\x69\x74\x69\x61\x6c\x69\x7a\x65\x20\x6f\x61\x75\x74\x68\x3a\x20\x27\x20\x2b\x20\x65\x72\x72\x6f\x72\x73\x29\x3b\x0a\x20\x20\x20\x20\x72\x65\x74\x75\x72\x6e\x3b\x0a\x20\x20\x7d\x0a\x0a\x20\x20\x24\x28\x27\x70\x72\x65\x20\x63\x6f\x64\x65\x27\x29\x2e\x65\x61\x63\x68\x28\x66\x75\x6e\x63\x74\x69\x6f\x6e\x20\x28\x69\x2c\x20\x65\x29\x20\x7b\x0a\x20\x20\x20\x20\x68\x6c\x6a\x73\x2e\x68\x69\x67\x68\x6c\x69\x67\x68\x74\x42\x6c\x6f\x63\x6b\x28\x65\x29\x0a\x20\x20\x7d\x29\x3b\x0a\x0a\x20\x20\x76\x61\x72\x20\x6f\x61\x75\x74\x68\x42\x74\x6e\x20\x3d\x20\x24\x28\x27\x2e\x61\x70\x69\x2d\x69\x63\x27\x29\x3b\x0a\x20\x20\x6f\x61\x75\x74\x68\x42\x74\x6e\x2e\x75\x6e\x62\x69\x6e\x64\x28\x29\x3b\x0a\x20\x20\x6f\x61\x75\x74\x68\x42\x74\x6e\x2e\x63\x6c\x69\x63\x6b\x28\x66\x75\x6e\x63\x74\x69\x6f\x6e\x20\x28\x73\x29\x20\x7b\x0a\x20\x20\x20\x20\x69\x66\x20\x28\x24\x28\x73\x2e\x74\x61\x72\x67\x65\x74\x29\x2e\x68\x61\x73\x43\x6c\x61\x73\x73\x28\x27\x62\x74\x6e\x2d\x64\x65\x66\x61\x75\x6c\x74\x27\x29\x29\x0a\x20\x20\x20\x20\x20\x20\x68\x61\x6e\x64\x6c\x65\x4c\x6f\x67\x69\x6e\x28\x29\x3b\x0a\x20\x20\x20\x20\x65\x6c\x73\x65\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x68\x61\x6e\x64\x6c\x65\x4c\x6f\x67\x6f\x75\x74\x28\x29\x3b\x0a\x20\x20\x20\x20\x7d\x0a\x20\x20\x7d\x29\x3b\x0a\x7d\x0a\x0a\x77\x69\x6e\x64\x6f\x77\x2e\x70\x72\x6f\x63\x65\x73\x73\x4f\x41\x75\x74\x68\x43\x6f\x64\x65\x20\x3d\x20\x66\x75\x6e\x63\x74\x69\x6f\x6e\x20\x70\x72\x6f\x63\x65\x73\x73\x4f\x41\x75\x74\x68\x43\x6f\x64\x65\x28\x64\x61\x74\x61\x29\x20\x7b\x0a\x20\x20\x76\x61\x72\x20\x70\x61\x72\x61\x6d\x73\x20\x3d\x20\x7b\x0a\x20\x20\x20\x20\x27\x63\x6c\x69\x65\x6e\x74\x5f\x69\x64\x27\x3a\x20\x63\x6c\x69\x65\x6e\x74\x49\x64\x2c\x0a\x20\x20\x20\x20\x27\x63\x6f\x64\x65\x27\x3a\x20\x64\x61\x74\x61\x2e\x63\x6f\x64\x65\x2c\x0a\x20\x20\x20\x20\x27\x67\x72\x61\x6e\x74\x5f\x74\x79\x70\x65\x27\x3a\x20\x27\x61\x75\x74\x68\x6f\x72\x69\x7a\x61\x74\x69\x6f\x6e\x5f\x63\x6f\x64\x65\x27\x2c\x0a\x20\x20\x20\x20\x27\x72\x65\x64\x69\x72\x65\x63\x74\x5f\x75\x72\x69\x27\x3a\x20\x72\x65\x64\x69\x72\x65\x63\x74\x5f\x75\x72\x69\x0a\x20\x20\x7d\x3b\x0a\x20\x20\x24\x2e\x61\x6a\x61\x78\x28\x0a\x20\x20\x20\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x75\x72\x6c\x3a\x20\x77\x69\x6e\x64\x6f\x77\x2e\x73\x77\x61\x67\x67\x65\x72\x55\x69\x2e\x74\x6f\x6b\x65\x6e\x55\x72\x6c\x2c\x0a\x20\x20\x20\x20\x20\x20\x74\x79\x70\x65\x3a\x20\x22\x50\x4f\x53\x54\x22\x2c\x0a\x20\x20\x20\x20\x20\x20\x64\x61\x74\x61\x3a\x20\x70\x61\x72\x61\x6d\x73\x2c\x0a\x20\x20\x20\x20\x20\x20\x73\x75\x63\x63\x65\x73\x73\x3a\x20\x66\x75\x6e\x63\x74\x69\x6f\x6e\x20\x28\x64\x61\x74\x61\x2c\x20\x74\x65\x78\x74\x53\x74\x61\x74\x75\x73\x2c\x20\x6a\x71\x58\x48\x52\x29\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x6f\x6e\x4f\x41\x75\x74\x68\x43\x6f\x6d\x70\x6c\x65\x74\x65\x28\x64\x61\x74\x61\x29\x3b\x0a\x20\x20\x20\x20\x20\x20\x7d\x2c\x0a\x20\x20\x20\x20\x20\x20\x65\x72\x72\x6f\x72\x3a\x20\x66\x75\x6e\x63\x74\x69\x6f\x6e\x20\x28\x6a\x71\x58\x48\x52\x2c\x20\x74\x65\x78\x74\x53\x74\x61\x74\x75\x73\x2c\x20\x65\x72\x72\x6f\x72\x54\x68\x72\x6f\x77\x6e\x29\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x6f\x6e\x4f\x41\x75\x74\x68\x43\x6f\x6d\x70\x6c\x65\x74\x65\x28\x22\x22\x29\x3b\x0a\x20\x20\x20\x20\x20\x20\x7d\x0a\x20\x20\x20\x20\x7d\x29\x3b\x0a\x7d\x3b\x0a\x0a\x77\x69\x6e\x64\x6f\x77\x2e\x6f\x6e\x4f\x41\x75\x74\x68\x43\x6f\x6d\x70\x6c\x65\x74\x65\x20\x3d\x20\x66\x75\x6e\x63\x74\x69\x6f\x6e\x20\x6f\x6e\x4f\x41\x75\x74\x68\x43\x6f\x6d\x70\x6c\x65\x74\x65\x28\x74\x6f\x6b\x65\x6e\x29\x20\x7b\x0a\x20\x20\x69\x66\x20\x28\x74\x6f\x6b\x65\x6e\x29\x20\x7b\x0a\x20\x20\x20\x20\x69\x66\x20\x28\x74\x6f\x6b\x65\x6e\x2e\x65\x72\x72\x6f\x72\x29\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x76\x61\x72\x20\x63\x68\x65\x63\x6b\x62\x6f\x78\x20\x3d\x20\x24\x28\x27\x69\x6e\x70\x75\x74\x5b\x74\x79\x70\x65\x3d\x63\x68\x65\x63\x6b\x62\x6f\x78\x5d\x2c\x2e\x73\x65\x63\x75\x72\x65\x64\x27\x29\x0a\x20\x20\x20\x20\x20\x20\x63\x68\x65\x63\x6b\x62\x6f\x78\x2e\x65\x61\x63\x68\x28\x66\x75\x6e\x63\x74\x69\x6f\x6e\x20\x28\x70\x6f\x73\x29\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x63\x68\x65\x63\x6b\x62\x6f\x78\x5b\x70\x6f\x73\x5d\x2e\x63\x68\x65\x63\x6b\x65\x64\x20\x3d\x20\x66\x61\x6c\x73\x65\x3b\x0a\x20\x20\x20\x20\x20\x20\x7d\x29\x3b\x0a\x20\x20\x20\x20\x20\x20\x61\x6c\x65\x72\x74\x28\x74\x6f\x6b\x65\x6e\x2e\x65\x72\x72\x6f\x72\x29\x3b\x0a\x20\x20\x20\x20\x7d\x0a\x20\x20\x20\x20\x65\x6c\x73\x65\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x76\x61\x72\x20\x62\x20\x3d\x20\x74\x6f\x6b\x65\x6e\x5b\x77\x69\x6e\x64\x6f\x77\x2e\x73\x77\x61\x67\x67\x65\x72\x55\x69\x2e\x74\x6f\x6b\x65\x6e\x4e\x61\x6d\x65\x5d\x3b\x0a\x20\x20\x20\x20\x20\x20\x69\x66\x20\x28\x62\x29\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x2f\x2f\x20\x69\x66\x20\x61\x6c\x6c\x20\x72\x6f\x6c\x65\x73\x20\x61\x72\x65\x20\x73\x61\x74\x69\x73\x66\x69\x65\x64\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x76\x61\x72\x20\x6f\x20\x3d\x20\x6e\x75\x6c\x6c\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x24\x2e\x65\x61\x63\x68\x28\x24\x28\x27\x2e\x61\x75\x74\x68\x20\x23\x61\x70\x69\x5f\x69\x6e\x66\x6f\x72\x6d\x61\x74\x69\x6f\x6e\x5f\x70\x61\x6e\x65\x6c\x27\x29\x2c\x20\x66\x75\x6e\x63\x74\x69\x6f\x6e\x20\x28\x6b\x2c\x20\x76\x29\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x76\x61\x72\x20\x63\x68\x69\x6c\x64\x72\x65\x6e\x20\x3d\x20\x76\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x69\x66\x20\x28\x63\x68\x69\x6c\x64\x72\x65\x6e\x20\x26\x26\x20\x63\x68\x69\x6c\x64\x72\x65\x6e\x2e\x63\x68\x69\x6c\x64\x4e\x6f\x64\x65\x73\x29\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x76\x61\x72\x20\x72\x65\x71\x75\x69\x72\x65\x64\x53\x63\x6f\x70\x65\x73\x20\x3d\x20\x5b\x5d\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x24\x2e\x65\x61\x63\x68\x28\x28\x63\x68\x69\x6c\x64\x72\x65\x6e\x2e\x63\x68\x69\x6c\x64\x4e\x6f\x64\x65\x73\x29\x2c\x20\x66\x75\x6e\x63\x74\x69\x6f\x6e\x20\x28\x6b\x31\x2c\x20\x76\x31\x29\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x76\x61\x72\x20\x69\x6e\x6e\x65\x72\x20\x3d\x20\x76\x31\x2e\x69\x6e\x6e\x65\x72\x48\x54\x4d\x4c\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x69\x66\x20\x28\x69\x6e\x6e\x65\x72\x29\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x72\x65\x71\x75\x69\x72\x65\x64\x53\x63\x6f\x70\x65\x73\x2e\x70\x75\x73\x68\x28\x69\x6e\x6e\x65\x72\x29\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x7d\x29\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x76\x61\x72\x20\x64\x69\x66\x66\x20\x3d\x20\x5b\x5d\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x66\x6f\x72\x20\x28\x76\x61\x72\x20\x69\x20\x3d\x20\x30\x3b\x20\x69\x20\x3c\x20\x72\x65\x71\x75\x69\x72\x65\x64\x53\x63\x6f\x70\x65\x73\x2e\x6c\x65\x6e\x67\x74\x68\x3b\x20\x69\x2b\x2b\x29\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x76\x61\x72\x20\x73\x20\x3d\x20\x72\x65\x71\x75\x69\x72\x65\x64\x53\x63\x6f\x70\x65\x73\x5b\x69\x5d\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x69\x66\x20\x28\x77\x69\x6e\x64\x6f\x77\x2e\x65\x6e\x61\x62\x6c\x65\x64\x53\x63\x6f\x70\x65\x73\x20\x26\x26\x20\x77\x69\x6e\x64\x6f\x77\x2e\x65\x6e\x61\x62\x6c\x65\x64\x53\x63\x6f\x70\x65\x73\x2e\x69\x6e\x64\x65\x78\x4f\x66\x28\x73\x29\x20\x3d\x3d\x20\x2d\x31\x29\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x64\x69\x66\x66\x2e\x70\x75\x73\x68\x28\x73\x29\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x7d\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x7d\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x69\x66\x20\x28\x64\x69\x66\x66\x2e\x6c\x65\x6e\x67\x74\x68\x20\x3e\x20\x30\x29\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x6f\x20\x3d\x20\x76\x2e\x70\x61\x72\x65\x6e\x74\x4e\x6f\x64\x65\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x2f\x2f\x20\x73\x6f\x72\x72\x79\x2c\x20\x6e\x6f\x74\x20\x61\x6c\x6c\x20\x73\x63\x6f\x70\x65\x73\x20\x61\x72\x65\x20\x73\x61\x74\x69\x73\x66\x69\x65\x64\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x24\x28\x6f\x29\x2e\x66\x69\x6e\x64\x28\x27\x2e\x61\x70\x69\x2d\x69\x63\x27\x29\x2e\x61\x64\x64\x43\x6c\x61\x73\x73\x28\x27\x62\x74\x6e\x2d\x77\x61\x72\x6e\x69\x6e\x67\x27\x29\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x24\x28\x6f\x29\x2e\x66\x69\x6e\x64\x28\x27\x2e\x61\x70\x69\x2d\x69\x63\x27\x29\x2e\x72\x65\x6d\x6f\x76\x65\x43\x6c\x61\x73\x73\x28\x27\x62\x74\x6e\x2d\x64\x65\x66\x61\x75\x6c\x74\x27\x29\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x24\x28\x6f\x29\x2e\x66\x69\x6e\x64\x28\x27\x2e\x61\x70\x69\x2d\x69\x63\x27\x29\x2e\x72\x65\x6d\x6f\x76\x65\x43\x6c\x61\x73\x73\x28\x27\x62\x74\x6e\x2d\x73\x75\x63\x63\x65\x73\x73\x27\x29\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x7d\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x65\x6c\x73\x65\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x6f\x20\x3d\x20\x76\x2e\x70\x61\x72\x65\x6e\x74\x4e\x6f\x64\x65\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x2f\x2f\x20\x61\x6c\x6c\x20\x73\x63\x6f\x70\x65\x73\x20\x61\x72\x65\x20\x73\x61\x74\x69\x73\x66\x69\x65\x64\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x24\x28\x6f\x29\x2e\x66\x69\x6e\x64\x28\x27\x2e\x61\x70\x69\x2d\x69\x63\x27\x29\x2e\x61\x64\x64\x43\x6c\x61\x73\x73\x28\x27\x62\x74\x6e\x2d\x73\x75\x63\x63\x65\x73\x73\x27\x29\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x24\x28\x6f\x29\x2e\x66\x69\x6e\x64\x28\x27\x2e\x61\x70\x69\x2d\x69\x63\x27\x29\x2e\x72\x65\x6d\x6f\x76\x65\x43\x6c\x61\x73\x73\x28\x27\x62\x74\x6e\x2d\x64\x65\x66\x61\x75\x6c\x74\x27\x29\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x24\x28\x6f\x29\x2e\x66\x69\x6e\x64\x28\x27\x2e\x61\x70\x69\x2d\x69\x63\x27\x29\x2e\x72\x65\x6d\x6f\x76\x65\x43\x6c\x61\x73\x73\x28\x27\x62\x74\x6e\x2d\x77\x61\x72\x6e\x69\x6e\x67\x27\x29\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x7d\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x7d\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x7d\x29\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x24\x28\x27\x2e\x61\x70\x69\x2d\x69\x63\x27\x29\x2e\x74\x65\x78\x74\x28\x27\x73\x69\x67\x6e\x6f\x75\x74\x27\x29\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x24\x28\x27\x23\x69\x6e\x70\x75\x74\x5f\x61\x70\x69\x4b\x65\x79\x27\x29\x2e\x76\x61\x6c\x28\x62\x29\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x77\x69\x6e\x64\x6f\x77\x2e\x73\x77\x61\x67\x67\x65\x72\x55\x69\x2e\x61\x70\x69\x2e\x63\x6c\x69\x65\x6e\x74\x41\x75\x74\x68\x6f\x72\x69\x7a\x61\x74\x69\x6f\x6e\x73\x2e\x61\x64\x64\x28\x6f\x61\x75\x74\x68\x32\x4b\x65\x79\x4e\x61\x6d\x65\x2c\x20\x6e\x65\x77\x20\x53\x77\x61\x67\x67\x65\x72\x43\x6c\x69\x65\x6e\x74\x2e\x41\x70\x69\x4b\x65\x79\x41\x75\x74\x68\x6f\x72\x69\x7a\x61\x74\x69\x6f\x6e\x28\x27\x41\x75\x74\x68\x6f\x72\x69\x7a\x61\x74\x69\x6f\x6e\x27\x2c\x20\x27\x42\x65\x61\x72\x65\x72\x20\x27\x20\x2b\x20\x62\x2c\x20\x27\x68\x65\x61\x64\x65\x72\x27\x29\x29\x3b\x0a\x20\x20\x20\x20\x20\x20\x7d\x0a\x20\x20\x20\x20\x7d\x0a\x20\x20\x7d\x0a\x7d\x3b\x0a")

func init() {

	f, err := FS.OpenFile(CTX, "/lib/swagger-oauth.js", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = f.Write(FileLibSwaggerOauthJs)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}
}
