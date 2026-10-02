---
page_title: "enable_ip_reputation"
subcategory: "Load Balancing"
description: "List of IP threat categories."
xcsh_docs: {"aliases": ["enable ip reputation"], "body_bytes": 2648, "body_sha256": "sha256:63d2fdccb4036086918abbf55bb51c0b9e4a6f8fda930e14d33b57f0b4fb6bbb", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:enable_ip_reputation", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:reference", "path": "documentation/data-sources/http_loadbalancer/properties/enable_ip_reputation/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2222220232120233-1130313100210223-1110323001300223-2031010002311213-3032120333310102-3022103310121023-0103023331030110-0220211011320211", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-018.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["enable_ip_reputation"], "schema_version": 1, "sections": [{"aliases": ["ip threat categories"], "anchor": "schema-enable_ip_reputation--ip_threat_categories", "description": "If the source IP matches on atleast one of the enabled IP threat categories, the request will be denied.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:enable_ip_reputation", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_ip_reputation", "ip_threat_categories"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/enable_ip_reputation/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "List of IP threat categories.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_ip_reputation

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- enable_ip_reputation

<a id="section"></a>

Type: `"single"`. Computed.

IP Threat Category List. List of IP threat categories.

Upstream description:

List of IP threat categories.

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

<a id="schema-enable_ip_reputation--ip_threat_categories"></a>

### ip_threat_categories property

Type: `["list", "string"]`. Computed.

\[Enum:
SPAM\_SOURCES|WINDOWS\_EXPLOITS|WEB\_ATTACKS|BOTNETS|SCANNERS|REPUTATION|PHISHING|PROXY|MOBILE\_THREATS|TOR\_PROXY|DENIAL\_OF\_SERVICE|NETWORK\]
If the source IP matches on atleast one of the enabled IP threat categories, the request will be
denied. Possible values are \`SPAM\_SOURCES\`, \`WINDOWS\_EXPLOITS\`, \`WEB\_ATTACKS\`, \`BOTNETS\`,
\`SCANNERS\`, \`REPUTATION\`, \`PHISHING\`, \`PROXY\`, \`MOBILE\_THREATS\`, \`TOR\_PROXY\`,
\`DENIAL\_OF\_SERVICE\`, \`NETWORK\`. Defaults to \`SPAM\_SOURCES\`.

Upstream description:

If the source IP matches on atleast one of the enabled IP threat categories, the request will be
denied.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
