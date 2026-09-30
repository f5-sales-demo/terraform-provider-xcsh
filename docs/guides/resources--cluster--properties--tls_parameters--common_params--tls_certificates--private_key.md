---
page_title: "tls_parameters.common_params.tls_certificates.private_key"
subcategory: ""
description: "tls_parameters.common_params.tls_certificates.private_key for xcsh_cluster."
xcsh_docs: {"aliases": [], "body_bytes": 2206, "body_sha256": "sha256:5924857cc71e85be453ec41cb71f1f2713c76fef31e5a974c6f79e4316cb5bac", "canonical_id": "xcsh-docs:resources:cluster:properties:tls_parameters:common_params:tls_certificates:private_key", "child_ids": ["xcsh-docs:resources:cluster:properties:tls_parameters:common_params:tls_certificates:private_key:blindfold_secret_info", "xcsh-docs:resources:cluster:properties:tls_parameters:common_params:tls_certificates:private_key:clear_secret_info"], "collection_id": "xcsh-docs:resources:cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:cluster:properties:tls_parameters:common_params:tls_certificates:private_key", "parent_id": "xcsh-docs:resources:cluster:properties:tls_parameters:common_params:tls_certificates", "path": "docs/guides/resources--cluster--properties--tls_parameters--common_params--tls_certificates--private_key.md", "provider_name": "cluster", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["tls_parameters", "common_params", "tls_certificates", "private_key"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cluster/properties/tls_parameters/common_params/tls_certificates/private_key/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "tls_parameters.common_params.tls_certificates.private_key for xcsh_cluster.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["clusterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# tls_parameters.common_params.tls_certificates.private_key

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md)
- [Property reference](resources--cluster--reference.md)
- [tls_parameters](resources--cluster--properties--tls_parameters.md)
- [tls_parameters.common_params](resources--cluster--properties--tls_parameters--common_params.md)
- [tls_parameters.common_params.tls_certificates](resources--cluster--properties--tls_parameters--common_params--tls_certificates.md)
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

- [blindfold_secret_info](resources--cluster--properties--tls_parameters--common_params--tls_certificates--private_key--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](resources--cluster--properties--tls_parameters--common_params--tls_certificates--private_key--clear_secret_info.md): complete subsection reference.

## Next pages

- [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info](resources--cluster--properties--tls_parameters--common_params--tls_certificates--private_key--blindfold_secret_info.md)
- [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info](resources--cluster--properties--tls_parameters--common_params--tls_certificates--private_key--clear_secret_info.md)
- [tls_parameters.common_params.tls_certificates](resources--cluster--properties--tls_parameters--common_params--tls_certificates.md)
- [xcsh_cluster](../resources/cluster.md)
