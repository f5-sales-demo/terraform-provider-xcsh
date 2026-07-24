// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

// Package docsterm provides terminology-normalisation transforms applied to
// generated Terraform provider documentation. It is deliberately a normal,
// testable package (not part of the //go:build ignore transform-docs.go tool)
// so its behaviour can be gated by `go test ./tools/...`.
package docsterm

import (
	"fmt"
	"regexp"
	"strings"
)

// FixUpstreamTerminology corrects upstream API terminology to pass textlint rules.
// Spelling corrections are handled by codespell --write-changes in the CI workflow.
func FixUpstreamTerminology(content string) string {
	// Protect markdown link URLs from terminology corrections.
	// Store URLs with placeholders, apply corrections, then restore.
	urlRegex := regexp.MustCompile(`\]\((https?://[^)]+)\)`)
	var savedURLs []string
	content = urlRegex.ReplaceAllStringFunc(content, func(match string) string {
		idx := len(savedURLs)
		savedURLs = append(savedURLs, match)
		return fmt.Sprintf("](##URL_%d##)", idx)
	})
	content = strings.NewReplacer(
		"User Name", "username",
		"Host Name", "hostname",
		"name space", "namespace",
		"Javascript", "JavaScript",
		"javascript", "JavaScript",
		"MAC OS", "macOS",
		"Clientside", "client-side",
		"Client Side", "client-side",
		"client side", "client-side",
		"server side", "server-side",
		"sub-class", "subclass",
		"Code Base", "codebase",
		"code base", "codebase",
		"Internet", "internet",
	).Replace(content)

	cdnRegex := regexp.MustCompile(`\bcdn\b`)
	content = cdnRegex.ReplaceAllString(content, "CDN")

	clickhouseRegex := regexp.MustCompile(`(?i)\bclickhouse\b`)
	content = clickhouseRegex.ReplaceAllStringFunc(content, func(_ string) string {
		return "ClickHouse"
	})

	sdkRegex := regexp.MustCompile(`\bSdk\b`)
	content = sdkRegex.ReplaceAllString(content, "SDK")

	githubRegex := regexp.MustCompile(`\b[Gg]ithub\b`)
	content = githubRegex.ReplaceAllString(content, "GitHub")

	gitlabRegex := regexp.MustCompile(`\b[Gg]itlab\b`)
	content = gitlabRegex.ReplaceAllString(content, "GitLab")

	bitbucketRegex := regexp.MustCompile(`\b[Bb]it[Bb]ucket\b`)
	content = bitbucketRegex.ReplaceAllString(content, "Bitbucket")

	dockerRegex := regexp.MustCompile(`\bdocker\b`)
	content = dockerRegex.ReplaceAllString(content, "Docker")

	ubuntuRegex := regexp.MustCompile(`\bubuntu\b`)
	content = ubuntuRegex.ReplaceAllString(content, "Ubuntu")

	azureRegex := regexp.MustCompile(`\bazure\b`)
	content = azureRegex.ReplaceAllString(content, "Azure")

	cassandraRegex := regexp.MustCompile(`\bcassandra\b`)
	content = cassandraRegex.ReplaceAllString(content, "Cassandra")

	mongodbRegex := regexp.MustCompile(`\bmongodb\b`)
	content = mongodbRegex.ReplaceAllString(content, "MongoDB")

	// NOTE: no Base64 -> base64 rewrite. Lowercasing prose "Base64" is not required
	// by any terminology rule, and doing so corrupted the real API enum token
	// "EncodingBase64" into "Encodingbase64" (Go regexp has no lookbehind to exclude
	// it). Preserve the token as authored in the schema.

	// Restore protected URLs
	for i, url := range savedURLs {
		content = strings.Replace(content, fmt.Sprintf("](##URL_%d##)", i), url, 1)
	}

	return content
}
