---
page_title: "ddos_profile.enable_ddos_mitigation"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["ddos profile enable ddos mitigation"], "body_bytes": 1243, "body_sha256": "sha256:99c7e0d0fbe926e393ece6a594f34c401ddea319c78dfafdf6fb509fae9cab99", "capabilities": ["dns"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_proxy:properties:ddos_profile:enable_ddos_mitigation", "parent_id": "xcsh-docs:resources:dns_proxy:properties:ddos_profile", "path": "documentation/resources/dns_proxy/properties/ddos_profile/enable_ddos_mitigation/index.md", "product": "distributed-cloud", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0113322033110122-3010012031011023-0120030310322113-1011300230111202-3001303212100313-3020111302222203-2333101200222321-0223001130313013", "registry_path": "docs/guides/resources--dns_proxy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ddos_profile", "enable_ddos_mitigation"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_proxy/properties/ddos_profile/enable_ddos_mitigation/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ddos_profile.enable_ddos_mitigation

Breadcrumbs:

- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/)
- [ddos_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/ddos_profile/)
- ddos_profile.enable_ddos_mitigation

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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
enable_ddos_mitigation = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [ddos_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/ddos_profile/)
- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/)
