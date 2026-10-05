---
page_title: "tls_parameters.common_params.tls_certificates.private_key"
subcategory: ""
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["tls parameters common params tls certificates private key"], "body_bytes": 2847, "body_sha256": "sha256:fd8e478f3ed216671c3191f32cbd25a9e4da64b97d3514d9e68887ad4b994eae", "capabilities": ["load-balancing.tls"], "category": null, "child_ids": ["xcsh-docs:resources:cluster:properties:tls_parameters:common_params:tls_certificates:private_key:blindfold_secret_info", "xcsh-docs:resources:cluster:properties:tls_parameters:common_params:tls_certificates:private_key:clear_secret_info"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:cluster:properties:tls_parameters:common_params:tls_certificates:private_key", "parent_id": "xcsh-docs:resources:cluster:properties:tls_parameters:common_params:tls_certificates", "path": "documentation/resources/cluster/properties/tls_parameters/common_params/tls_certificates/private_key/index.md", "product": "distributed-cloud", "provider_name": "cluster", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-0112322121010113-2232031323011101-1331110203332032-1103021020010321-1303003331100031-1100121002230200-0112002301111132-2220220321333331", "registry_path": "docs/guides/resources--cluster--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "tls_parameters.common_params.tls_certificates.private_key:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:tls_parameters:common_params:tls_certificates:private_key:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_parameters.common_params.tls_certificates.private_key:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:tls_parameters:common_params:tls_certificates:private_key:clear_secret_info", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["tls_parameters", "common_params", "tls_certificates", "private_key"], "schema_version": 1, "sections": [{"aliases": ["tls parameters common params tls certificates private key blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:resources:cluster:properties:tls_parameters:common_params:tls_certificates:private_key:blindfold_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-tls_parameters--common_params--tls_certificates--private_key--blindfold_secret_info--location", "enforcement": "provider-schema", "group": "tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info:RequiredObjectAttributes:location", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:tls_parameters:common_params:tls_certificates:private_key:blindfold_secret_info", "type": "requires"}], "schema_path": ["tls_parameters", "common_params", "tls_certificates", "private_key", "blindfold_secret_info"], "syntax": "block", "type": "object"}, {"aliases": ["tls parameters common params tls certificates private key clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:resources:cluster:properties:tls_parameters:common_params:tls_certificates:private_key:clear_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-tls_parameters--common_params--tls_certificates--private_key--clear_secret_info--url", "enforcement": "provider-schema", "group": "tls_parameters.common_params.tls_certificates.private_key.clear_secret_info:RequiredObjectAttributes:url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:tls_parameters:common_params:tls_certificates:private_key:clear_secret_info", "type": "requires"}], "schema_path": ["tls_parameters", "common_params", "tls_certificates", "private_key", "clear_secret_info"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cluster/properties/tls_parameters/common_params/tls_certificates/private_key/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["clusterCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_parameters.common_params.tls_certificates.private_key

Breadcrumbs:

- [xcsh_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/)
- [tls_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/)
- [tls_parameters.common_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/common_params/)
- [tls_parameters.common_params.tls_certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/common_params/tls_certificates/)
- tls_parameters.common_params.tls_certificates.private_key

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
private_key {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/common_params/tls_certificates/private_key/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/common_params/tls_certificates/private_key/clear_secret_info/): complete subsection reference.

## Next pages

- [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/common_params/tls_certificates/private_key/blindfold_secret_info/)
- [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/common_params/tls_certificates/private_key/clear_secret_info/)
- [tls_parameters.common_params.tls_certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/common_params/tls_certificates/)
- [xcsh_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/)
