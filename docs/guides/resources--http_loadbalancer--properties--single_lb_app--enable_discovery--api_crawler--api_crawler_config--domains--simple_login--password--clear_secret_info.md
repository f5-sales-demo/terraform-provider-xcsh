---
page_title: "single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info"
subcategory: "Load Balancing"
description: "single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 4868, "body_sha256": "sha256:730e6ea67f1f4252db89b999638e8a6bb3fb30e47039bc4d50987c9e04c5a41f", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_crawler:api_crawler_config:domains:simple_login:password:clear_secret_info", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_crawler:api_crawler_config:domains:simple_login:password:clear_secret_info", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_crawler:api_crawler_config:domains:simple_login:password", "path": "docs/guides/resources--http_loadbalancer--properties--single_lb_app--enable_discovery--api_crawler--api_crawler_config--domains--simple_login--password--clear_secret_info.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["single_lb_app", "enable_discovery", "api_crawler", "api_crawler_config", "domains", "simple_login", "password", "clear_secret_info"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_crawler/api_crawler_config/domains/simple_login/password/clear_secret_info/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [single_lb_app](resources--http_loadbalancer--properties--single_lb_app.md)
- [single_lb_app.enable_discovery](resources--http_loadbalancer--properties--single_lb_app--enable_discovery.md)
- [single_lb_app.enable_discovery.api_crawler](resources--http_loadbalancer--properties--single_lb_app--enable_discovery--api_crawler.md)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config](resources--http_loadbalancer--properties--single_lb_app--enable_discovery--api_crawler--api_crawler_config.md)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains](resources--http_loadbalancer--properties--single_lb_app--enable_discovery--api_crawler--api_crawler_config--domains.md)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login](resources--http_loadbalancer--properties--single_lb_app--enable_discovery--api_crawler--api_crawler_config--domains--simple_login.md)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password](resources--http_loadbalancer--properties--single_lb_app--enable_discovery--api_crawler--api_crawler_config--domains--simple_login--password.md)
- single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
clear_secret_info {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-single_lb_app--enable_discovery--api_crawler--api_crawler_config--domains--simple_login--password--clear_secret_info--provider_ref"></a>

### provider_ref property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="schema-single_lb_app--enable_discovery--api_crawler--api_crawler_config--domains--simple_login--password--clear_secret_info--url"></a>

### url property

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

## Next pages

- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password](resources--http_loadbalancer--properties--single_lb_app--enable_discovery--api_crawler--api_crawler_config--domains--simple_login--password.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
