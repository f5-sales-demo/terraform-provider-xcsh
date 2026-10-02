---
page_title: "discovery_k8s.publish_info.dns_delegation"
subcategory: ""
description: "Configuration parameter for dns delegation."
xcsh_docs: {"aliases": ["discovery k8s publish info dns delegation"], "body_bytes": 2942, "body_sha256": "sha256:282d51f41e26081629845daa9a65f7d0d1387ff82d96ae925b618fa7fb577434", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:discovery:properties:discovery_k8s:publish_info:dns_delegation", "parent_id": "xcsh-docs:data-sources:discovery:properties:discovery_k8s:publish_info", "path": "documentation/data-sources/discovery/properties/discovery_k8s/publish_info/dns_delegation/index.md", "product": "distributed-cloud", "provider_name": "discovery", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0021110011333022-3031323120010221-3331021222110233-2032301330012323-0021101112313013-0221323000310213-0010322112333203-2312212121122233", "registry_path": "docs/guides/data-sources--discovery--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["discovery_k8s", "publish_info", "dns_delegation"], "schema_version": 1, "sections": [{"aliases": ["dns mode"], "anchor": "schema-discovery_k8s--publish_info--dns_delegation--dns_mode", "description": "Two modes are possible CoreDNS: Whether external K8s cluster is running core-DNS KubeDNS: External K8s cluster is running kube-DNS.", "document_id": "xcsh-docs:data-sources:discovery:properties:discovery_k8s:publish_info:dns_delegation", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["discovery_k8s", "publish_info", "dns_delegation", "dns_mode"], "syntax": "attribute", "type": "string"}, {"aliases": ["subdomain"], "anchor": "schema-discovery_k8s--publish_info--dns_delegation--subdomain", "description": "The DNS subdomain for which F5XC will respond to DNS queries.", "document_id": "xcsh-docs:data-sources:discovery:properties:discovery_k8s:publish_info:dns_delegation", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["discovery_k8s", "publish_info", "dns_delegation", "subdomain"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/discovery/properties/discovery_k8s/publish_info/dns_delegation/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Configuration parameter for dns delegation.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["discoveryCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# discovery_k8s.publish_info.dns_delegation

Breadcrumbs:

- [xcsh_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/)
- [discovery_k8s](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_k8s/)
- [discovery_k8s.publish_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_k8s/publish_info/)
- discovery_k8s.publish_info.dns_delegation

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for dns delegation.

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

## Direct properties

<a id="schema-discovery_k8s--publish_info--dns_delegation--dns_mode"></a>

### dns_mode property

Type: `"string"`. Computed.

\[Enum: CORE\_DNS|KUBE\_DNS\] Two modes are possible CoreDNS: Whether external K8s cluster is
running core-DNS KubeDNS: External K8s cluster is running kube-DNS. Possible values are
\`CORE\_DNS\`, \`KUBE\_DNS\`. Defaults to \`CORE\_DNS\`.

Upstream description:

Two modes are possible

CoreDNS: Whether external K8s cluster is running core-DNS KubeDNS: External K8s cluster is running
kube-DNS.

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

Type: `"string"`. Computed.

The DNS subdomain for which F5XC will respond to DNS queries.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

## Next pages

- [discovery_k8s.publish_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_k8s/publish_info/)
- [xcsh_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/)
