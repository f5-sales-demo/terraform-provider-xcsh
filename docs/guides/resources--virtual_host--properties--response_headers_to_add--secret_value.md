---
page_title: "response_headers_to_add.secret_value"
subcategory: ""
description: "response_headers_to_add.secret_value for xcsh_virtual_host."
xcsh_docs: {"aliases": [], "body_bytes": 1825, "body_sha256": "sha256:1836c2762e83926929a5175bc44241bedf6cdd121ae7b1797c4693f20c3e2cc4", "canonical_id": "xcsh-docs:resources:virtual_host:properties:response_headers_to_add:secret_value", "child_ids": ["xcsh-docs:resources:virtual_host:properties:response_headers_to_add:secret_value:blindfold_secret_info", "xcsh-docs:resources:virtual_host:properties:response_headers_to_add:secret_value:clear_secret_info"], "collection_id": "xcsh-docs:resources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_host:properties:response_headers_to_add:secret_value", "parent_id": "xcsh-docs:resources:virtual_host:properties:response_headers_to_add", "path": "docs/guides/resources--virtual_host--properties--response_headers_to_add--secret_value.md", "provider_name": "virtual_host", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["response_headers_to_add", "secret_value"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_host/properties/response_headers_to_add/secret_value/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "response_headers_to_add.secret_value for xcsh_virtual_host.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# response_headers_to_add.secret_value

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md)
- [Property reference](resources--virtual_host--reference.md)
- [response_headers_to_add](resources--virtual_host--properties--response_headers_to_add.md)
- response_headers_to_add.secret_value

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
secret_value {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](resources--virtual_host--properties--response_headers_to_add--secret_value--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](resources--virtual_host--properties--response_headers_to_add--secret_value--clear_secret_info.md): complete subsection reference.

## Next pages

- [response_headers_to_add.secret_value.blindfold_secret_info](resources--virtual_host--properties--response_headers_to_add--secret_value--blindfold_secret_info.md)
- [response_headers_to_add.secret_value.clear_secret_info](resources--virtual_host--properties--response_headers_to_add--secret_value--clear_secret_info.md)
- [response_headers_to_add](resources--virtual_host--properties--response_headers_to_add.md)
- [xcsh_virtual_host](../resources/virtual_host.md)
