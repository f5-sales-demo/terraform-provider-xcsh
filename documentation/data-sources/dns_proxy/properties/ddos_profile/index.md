---
page_title: "ddos_profile"
subcategory: ""
description: "DDoS Protection Rule for DNS."
xcsh_docs: {"aliases": ["ddos profile"], "body_bytes": 1757, "body_sha256": "sha256:5c811dc0856f5e688e50811b1cd104495a927d6aa348f662994aee59580777d9", "capabilities": ["dns"], "category": "dns", "child_ids": ["xcsh-docs:data-sources:dns_proxy:properties:ddos_profile:disable_ddos_mitigation", "xcsh-docs:data-sources:dns_proxy:properties:ddos_profile:enable_ddos_mitigation"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_proxy:properties:ddos_profile", "parent_id": "xcsh-docs:data-sources:dns_proxy:reference", "path": "documentation/data-sources/dns_proxy/properties/ddos_profile/index.md", "product": "distributed-cloud", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0232013023122233-1012213323101023-2001112012330030-3013030103202023-2011131011012030-2021102202320120-0100103221220130-0320000200033100", "registry_path": "docs/guides/data-sources--dns_proxy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ddos_profile"], "schema_version": 1, "sections": [{"aliases": ["disable ddos mitigation"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:dns_proxy:properties:ddos_profile:disable_ddos_mitigation", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ddos_profile", "disable_ddos_mitigation"], "syntax": "attribute", "type": "object"}, {"aliases": ["enable ddos mitigation"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:dns_proxy:properties:ddos_profile:enable_ddos_mitigation", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ddos_profile", "enable_ddos_mitigation"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_proxy/properties/ddos_profile/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "DDoS Protection Rule for DNS.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ddos_profile

Breadcrumbs:

- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/)
- ddos_profile

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for ddos profile.

Upstream description:

DDoS Protection Rule for DNS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ddos_mitigation_choice": "[\"disable_ddos_mitigation\",\"enable_ddos_mitigation\"]"
}
```

## Direct properties

- [disable_ddos_mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/ddos_profile/disable_ddos_mitigation/): complete subsection reference.

- [enable_ddos_mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/ddos_profile/enable_ddos_mitigation/): complete subsection reference.

## Next pages

- [ddos_profile.disable_ddos_mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/ddos_profile/disable_ddos_mitigation/)
- [ddos_profile.enable_ddos_mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/ddos_profile/enable_ddos_mitigation/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/)
- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/)
