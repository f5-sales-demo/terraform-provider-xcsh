---
page_title: "default_sriov_interface"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["default sriov interface"], "body_bytes": 1650, "body_sha256": "sha256:b1fe7e176d5b46796e2f7c888f9f7b525e38ebc74a1de90b0b7f1b9cbcc98e2e", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:default_sriov_interface", "parent_id": "xcsh-docs:resources:voltstack_site:reference", "path": "documentation/resources/voltstack_site/properties/default_sriov_interface/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0333212220131232-2232130221010210-1121201121210113-3310332003011111-3130113303011212-3232313102100302-3300112310202231-0003110030333311", "registry_path": "docs/guides/resources--voltstack_site--reference--group-009.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["default_sriov_interface"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/default_sriov_interface/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_sriov_interface

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/)
- default_sriov_interface

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: default\_sriov\_interface, sriov\_interfaces; Default: default\_sriov\_interface\]
Configuration parameter for default sriov interface.

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

OneOf alternatives in this subsection:

- [default_sriov_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/default_sriov_interface/#section)
- [sriov_interfaces](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/sriov_interfaces/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
default_sriov_interface = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
