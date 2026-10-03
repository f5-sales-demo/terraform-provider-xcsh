---
page_title: "infra.internet_proxy"
subcategory: ""
description: "Proxy describes OPTIONS for HTTP or HTTPS proxy configurations."
xcsh_docs: {"aliases": ["infra internet proxy"], "body_bytes": 4676, "body_sha256": "sha256:df28b3b4a84281c512803d504d23e958c0b14e41af26af9e230a3570da2bd22d", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:registration:collection", "completeness": "complete", "id": "xcsh-docs:resources:registration:properties:infra:internet_proxy", "parent_id": "xcsh-docs:resources:registration:properties:infra", "path": "documentation/resources/registration/properties/infra/internet_proxy/index.md", "product": "distributed-cloud", "provider_name": "registration", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0333232120213313-2131312201312220-0332010123213011-1100232332221223-3130030323031202-2201003302332033-3110013202300010-1101212310212223", "registry_path": "docs/guides/resources--registration--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["infra", "internet_proxy"], "schema_version": 1, "sections": [{"aliases": ["infra internet proxy http proxy"], "anchor": "schema-infra--internet_proxy--http_proxy", "description": "It will be used as the proxy URL for HTTP requests and HTTPS requests unless overridden by HTTPSProxy or NoProxy.", "document_id": "xcsh-docs:resources:registration:properties:infra:internet_proxy", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "internet_proxy", "http_proxy"], "syntax": "attribute", "type": "string"}, {"aliases": ["infra internet proxy https proxy"], "anchor": "schema-infra--internet_proxy--https_proxy", "description": "It will be used as the proxy URL for HTTPS requests unless overridden by NoProxy.", "document_id": "xcsh-docs:resources:registration:properties:infra:internet_proxy", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "internet_proxy", "https_proxy"], "syntax": "attribute", "type": "string"}, {"aliases": ["infra internet proxy no proxy"], "anchor": "schema-infra--internet_proxy--no_proxy", "description": "It specifies a string that contains comma-separated values specifying hosts that should be excluded from proxying. Each value is represented by an IP address prefix (192.0.2.103), an IP address prefix in CIDR notation (192.0.2.103/8), a domain name, or a special DNS label (*). An IP address prefix and domain name can", "document_id": "xcsh-docs:resources:registration:properties:infra:internet_proxy", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "internet_proxy", "no_proxy"], "syntax": "attribute", "type": "string"}, {"aliases": ["infra internet proxy proxy cacert url"], "anchor": "schema-infra--internet_proxy--proxy_cacert_url", "description": "Allow optional different trust-store for proxy in HTTP CONNECT step by picking proxy CA certificate value.", "document_id": "xcsh-docs:resources:registration:properties:infra:internet_proxy", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "internet_proxy", "proxy_cacert_url"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/registration/properties/infra/internet_proxy/index.txt", "spec_pin_digest": "sha256:0e278b4afb598ac8628ac74dca6a1b2831cd8607de1d26f56e72e3461a7440c7", "summary": "Proxy describes OPTIONS for HTTP or HTTPS proxy configurations.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.0", "schema_components": ["registrationCreateRequest"], "target_commit": "ac024ccbfb8b9f84412813f3e2ab9f5821937ef2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# infra.internet_proxy

Breadcrumbs:

- [xcsh_registration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/)
- [infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/infra/)
- infra.internet_proxy

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Proxy describes OPTIONS for HTTP or HTTPS proxy configurations.

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
internet_proxy {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-infra--internet_proxy--http_proxy"></a>

### http_proxy property

Type: `"string"`. Optional.

It will be used as the proxy URL for HTTP requests and HTTPS requests unless overridden by
HTTPSProxy or NoProxy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-infra--internet_proxy--https_proxy"></a>

### https_proxy property

Type: `"string"`. Optional.

It will be used as the proxy URL for HTTPS requests unless overridden by NoProxy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-infra--internet_proxy--no_proxy"></a>

### no_proxy property

Type: `"string"`. Optional.

It specifies a string that contains comma-separated values specifying hosts that should be excluded
from proxying. Each value is represented by an IP address prefix (192.0.2.103), an IP address prefix
in CIDR notation (192.0.2.103/8), a domain name, or a special DNS label (\*). An IP address..

Upstream description:

It specifies a string that contains comma-separated values specifying hosts that should be excluded
from proxying. Each value is represented by an IP address prefix (192.0.2.103), an IP address prefix
in CIDR notation (192.0.2.103/8), a domain name, or a special DNS label (\*). An IP address prefix
and domain name can also include a literal port number (192.0.2.103:80). A domain name matches that
name and all subdomains. A domain name with a leading "." matches subdomains only. For example
"example.com" matches "example.com" and "bar.example.com"; ".y.com" matches "x.y.com" but not
"y.com". A single asterisk (\*) indicates that no proxying should be done.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-infra--internet_proxy--proxy_cacert_url"></a>

### proxy_cacert_url property

Type: `"string"`. Optional.

Allow optional different trust-store for proxy in HTTP CONNECT step by picking proxy CA certificate
value.

Upstream description:

Allow optional different trust-store for proxy in HTTP CONNECT step by picking proxy CA certificate
value.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## Next pages

- [infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/infra/)
- [xcsh_registration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/)
