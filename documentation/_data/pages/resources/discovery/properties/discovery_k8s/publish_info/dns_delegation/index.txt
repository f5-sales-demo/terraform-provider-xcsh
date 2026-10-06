---
page_title: "discovery_k8s.publish_info.dns_delegation"
subcategory: ""
description: "Configuration parameter for dns delegation."
xcsh_docs: {"aliases": ["discovery k8s publish info dns delegation"], "body_bytes": 3070, "body_sha256": "sha256:e67ab9992ce7b36ae12dcff932eec6e446aea5e6efdd5a3af8f0621b966630f5", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:resources:discovery:properties:discovery_k8s:publish_info:dns_delegation", "parent_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:publish_info", "path": "documentation/resources/discovery/properties/discovery_k8s/publish_info/dns_delegation/index.md", "product": "distributed-cloud", "provider_name": "discovery", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-0213033233320133-2313331011023030-3300221322303212-3012012303031332-2322003103212311-2232013223120231-0223002211302120-2121012230012320", "registry_path": "docs/guides/resources--discovery--reference--group-002.md", "relationships": [{"anchor": "schema-discovery_k8s--publish_info--dns_delegation--subdomain", "enforcement": "provider-schema", "group": "discovery_k8s.publish_info.dns_delegation:RequiredObjectAttributes:subdomain", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:publish_info:dns_delegation", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["discovery_k8s", "publish_info", "dns_delegation"], "schema_version": 1, "sections": [{"aliases": ["discovery k8s publish info dns delegation dns mode"], "anchor": "schema-discovery_k8s--publish_info--dns_delegation--dns_mode", "description": "Two modes are possible CoreDNS: Whether external K8s cluster is running core-DNS KubeDNS: External K8s cluster is running kube-DNS.", "document_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:publish_info:dns_delegation", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["discovery_k8s", "publish_info", "dns_delegation", "dns_mode"], "syntax": "attribute", "type": "string"}, {"aliases": ["discovery k8s publish info dns delegation subdomain"], "anchor": "schema-discovery_k8s--publish_info--dns_delegation--subdomain", "description": "The DNS subdomain for which F5XC will respond to DNS queries.", "document_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:publish_info:dns_delegation", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["discovery_k8s", "publish_info", "dns_delegation", "subdomain"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/discovery/properties/discovery_k8s/publish_info/dns_delegation/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Configuration parameter for dns delegation.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["discoveryCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# discovery_k8s.publish_info.dns_delegation

Breadcrumbs:

- [xcsh_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/)
- [discovery_k8s](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/)
- [discovery_k8s.publish_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/publish_info/)
- discovery_k8s.publish_info.dns_delegation

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for dns delegation.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("subdomain")}
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
dns_delegation {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-discovery_k8s--publish_info--dns_delegation--dns_mode"></a>

### dns_mode property

Type: `"string"`. Optional.

\[Enum: CORE\_DNS|KUBE\_DNS\] Two modes are possible CoreDNS: Whether external K8s cluster is
running core-DNS KubeDNS: External K8s cluster is running kube-DNS. Possible values are
\`CORE\_DNS\`, \`KUBE\_DNS\`. Defaults to \`CORE\_DNS\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("CORE_DNS",
    "KUBE_DNS"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "CORE_DNS",
  "enum": [
    "CORE_DNS",
    "KUBE_DNS"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-discovery_k8s--publish_info--dns_delegation--subdomain"></a>

### subdomain property

Type: `"string"`. Optional.

The DNS subdomain for which F5XC will respond to DNS queries.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```
