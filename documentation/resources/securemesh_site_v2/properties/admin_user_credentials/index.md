---
page_title: "admin_user_credentials"
subcategory: ""
description: "Setup user credentials to manage access to nodes belonging to the site. When configured, 'admin' user will be setup and customers can access these nodes via either the node local WebUI or via SSH to access shell/CLI Ensure 'Node Local Services' are enabled to allow for required access."
xcsh_docs: {"aliases": ["admin user credentials", "authentication", "credential setup", "credentials"], "body_bytes": 3257, "body_sha256": "sha256:cebc685ae4853040c6386d43bc8a53f82f627306e16588010044f4c5946ed549", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:admin_user_credentials:admin_password"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:admin_user_credentials", "parent_id": "xcsh-docs:resources:securemesh_site_v2:reference", "path": "documentation/resources/securemesh_site_v2/properties/admin_user_credentials/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-2323112033210020-2113012223203221-0322233010022032-2122320020101010-2102301101033021-3002303312010131-2012011322012232-0231103113201210", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["admin_user_credentials"], "schema_version": 1, "sections": [{"aliases": ["admin user credentials admin password"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:admin_user_credentials:admin_password", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "admin_user_credentials.admin_password:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:admin_user_credentials:admin_password:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "admin_user_credentials.admin_password:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:admin_user_credentials:admin_password:clear_secret_info", "type": "conflicts"}], "schema_path": ["admin_user_credentials", "admin_password"], "syntax": "block", "type": "object"}, {"aliases": ["admin user credentials ssh key"], "anchor": "schema-admin_user_credentials--ssh_key", "description": "Provided Public SSH key can be used for accessing nodes of the site. When provided, customers can SSH to the nodes of this Customer Edge site using admin as the user.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:admin_user_credentials", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["admin_user_credentials", "ssh_key"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/admin_user_credentials/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Setup user credentials to manage access to nodes belonging to the site. When configured, 'admin' user will be setup and customers can access these nodes via either the node local WebUI or via SSH to access shell/CLI Ensure 'Node Local Services' are enabled to allow for required access.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# admin_user_credentials

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- admin_user_credentials

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Setup user credentials to manage access to nodes belonging to the site. When configured, 'admin'
user will be setup and customers can access these nodes via either the node local WebUI or via SSH
to access shell/CLI Ensure 'Node Local Services' are enabled to allow for required access.

Upstream description:

Setup user credentials to manage access to nodes belonging to the site. When configured, 'admin'
user will be setup and customers can access these nodes via either the node local WebUI or via SSH
to access shell/CLI Ensure 'Node Local Services' are enabled to allow for required access.

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
admin_user_credentials {
  # Configure direct properties listed below.
}
```

## Direct properties

- [admin_password](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/admin_user_credentials/admin_password/): complete subsection reference.

<a id="schema-admin_user_credentials--ssh_key"></a>

### ssh_key property

Type: `"string"`. Optional.

Provided Public SSH key can be used for accessing nodes of the site. When provided, customers can
SSH to the nodes of this Customer Edge site using admin as the user.

Upstream description:

Provided Public SSH key can be used for accessing nodes of the site. When provided, customers can
SSH to the nodes of this Customer Edge site using admin as the user.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(8192),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8192,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8192,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8192"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8192"
  }
}
```

## Next pages

- [admin_user_credentials.admin_password](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/admin_user_credentials/admin_password/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
