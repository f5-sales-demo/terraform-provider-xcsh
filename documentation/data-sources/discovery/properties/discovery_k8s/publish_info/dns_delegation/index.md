---
page_title: "discovery_k8s.publish_info.dns_delegation"
subcategory: ""
description: "Configuration parameter for dns delegation."
xcsh_docs: {"aliases": ["discovery k8s publish info dns delegation"], "body_bytes": 2519, "body_sha256": "sha256:8c468c321132fa3419eb2c9ab93273042f20a11918d4640306bf05593bbd4446", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:discovery:properties:discovery_k8s:publish_info:dns_delegation", "parent_id": "xcsh-docs:data-sources:discovery:properties:discovery_k8s:publish_info", "path": "documentation/data-sources/discovery/properties/discovery_k8s/publish_info/dns_delegation/index.md", "product": "distributed-cloud", "provider_name": "discovery", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-0021110011333022-3031323120010221-3331021222110233-2032301330012323-0021101112313013-0221323000310213-0010322112333203-2312212121122233", "registry_path": "docs/guides/data-sources--discovery--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["discovery_k8s", "publish_info", "dns_delegation"], "schema_version": 1, "sections": [{"aliases": ["discovery k8s publish info dns delegation dns mode"], "anchor": "schema-discovery_k8s--publish_info--dns_delegation--dns_mode", "description": "Two modes are possible CoreDNS: Whether external K8s cluster is running core-DNS KubeDNS: External K8s cluster is running kube-DNS.", "document_id": "xcsh-docs:data-sources:discovery:properties:discovery_k8s:publish_info:dns_delegation", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["discovery_k8s", "publish_info", "dns_delegation", "dns_mode"], "syntax": "attribute", "type": "string"}, {"aliases": ["discovery k8s publish info dns delegation subdomain"], "anchor": "schema-discovery_k8s--publish_info--dns_delegation--subdomain", "description": "The DNS subdomain for which F5XC will respond to DNS queries.", "document_id": "xcsh-docs:data-sources:discovery:properties:discovery_k8s:publish_info:dns_delegation", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["discovery_k8s", "publish_info", "dns_delegation", "subdomain"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/discovery/properties/discovery_k8s/publish_info/dns_delegation/index.txt", "spec_pin_digest": "sha256:ea0d80973b5efe05eb3575be6beda9c96b86701d264dc1aee3972e6a23ebb3fb", "summary": "Configuration parameter for dns delegation.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.3", "schema_components": ["discoveryCreateRequest"], "target_commit": "6ee1e85b2e917d58078f38a270b790077b3fba12"}}
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
      "validatedAt": "2026-10-17T10:10:42+00:00"
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
