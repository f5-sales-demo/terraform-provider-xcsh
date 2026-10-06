---
page_title: "ipsec.ike_parameters.use_default_remote_ike_id"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["ipsec ike parameters use default remote ike id"], "body_bytes": 1187, "body_sha256": "sha256:6b479cf7e28162e801c063d81bd454cabf83b6cf261e08cbf35c5e09d90b0612", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:external_connector:collection", "completeness": "complete", "id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:use_default_remote_ike_id", "parent_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters", "path": "documentation/resources/external_connector/properties/ipsec/ike_parameters/use_default_remote_ike_id/index.md", "product": "distributed-cloud", "provider_name": "external_connector", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-2332221223221312-2330120323211321-3330111123031100-3130000121101013-2201300311311122-2332203102103332-2032230023131002-0213122010103222", "registry_path": "docs/guides/resources--external_connector--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ipsec", "ike_parameters", "use_default_remote_ike_id"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/external_connector/properties/ipsec/ike_parameters/use_default_remote_ike_id/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["external_connectorCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ipsec.ike_parameters.use_default_remote_ike_id

Breadcrumbs:

- [xcsh_external_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/)
- [ipsec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/)
- [ipsec.ike_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/ipsec/ike_parameters/)
- ipsec.ike_parameters.use_default_remote_ike_id

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

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

Terraform syntax:

```terraform
use_default_remote_ike_id = {}
```

This is an empty object or choice marker. It has no direct properties.
