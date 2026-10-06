---
page_title: "dualstack"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["dualstack"], "body_bytes": 1336, "body_sha256": "sha256:215082c50b8fdc385e9223c7757a587bae16c58a0d83bcc90f5a431777860bb3", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:advertise_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:advertise_policy:properties:dualstack", "parent_id": "xcsh-docs:resources:advertise_policy:reference", "path": "documentation/resources/advertise_policy/properties/dualstack/index.md", "product": "distributed-cloud", "provider_name": "advertise_policy", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-3130123120101222-1220203221123220-1312323010232302-0132230022313313-3223133102023311-2231101311030311-0233132213000003-3213202311001321", "registry_path": "docs/guides/resources--advertise_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["dualstack"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/advertise_policy/properties/dualstack/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["advertise_policyCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# dualstack

Breadcrumbs:

- [xcsh_advertise_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/advertise_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/advertise_policy/properties/)
- dualstack

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: dualstack, ipv4, ipv6\] Enable this option

Additional upstream details:

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

- [dualstack](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/advertise_policy/properties/dualstack/#section)
- [ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/advertise_policy/properties/ipv4/#section)
- [ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/advertise_policy/properties/ipv6/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
dualstack = {}
```

This is an empty object or choice marker. It has no direct properties.
