---
page_title: "discovery_consul.access_info.http_basic_auth_info.passwd_url"
subcategory: ""
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["discovery consul access info http basic auth info passwd url"], "body_bytes": 2916, "body_sha256": "sha256:dbe478150fe749f512de0cd48aa9d0222e668fa866940117c5a3435675026a7e", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:discovery:properties:discovery_consul:access_info:http_basic_auth_info:passwd_url:blindfold_secret_info", "xcsh-docs:resources:discovery:properties:discovery_consul:access_info:http_basic_auth_info:passwd_url:clear_secret_info"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:resources:discovery:properties:discovery_consul:access_info:http_basic_auth_info:passwd_url", "parent_id": "xcsh-docs:resources:discovery:properties:discovery_consul:access_info:http_basic_auth_info", "path": "documentation/resources/discovery/properties/discovery_consul/access_info/http_basic_auth_info/passwd_url/index.md", "product": "distributed-cloud", "provider_name": "discovery", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1120100130133100-3333202313112000-2201131010110123-3312220211222320-3230132002122323-1232211201113033-1112032210023213-3230323120221220", "registry_path": "docs/guides/resources--discovery--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "discovery_consul.access_info.http_basic_auth_info.passwd_url:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:discovery_consul:access_info:http_basic_auth_info:passwd_url:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "discovery_consul.access_info.http_basic_auth_info.passwd_url:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:discovery_consul:access_info:http_basic_auth_info:passwd_url:clear_secret_info", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["discovery_consul", "access_info", "http_basic_auth_info", "passwd_url"], "schema_version": 1, "sections": [{"aliases": ["blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:resources:discovery:properties:discovery_consul:access_info:http_basic_auth_info:passwd_url:blindfold_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-discovery_consul--access_info--http_basic_auth_info--passwd_url--blindfold_secret_info--location", "enforcement": "provider-schema", "group": "discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_info:RequiredObjectAttributes:location", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:discovery_consul:access_info:http_basic_auth_info:passwd_url:blindfold_secret_info", "type": "requires"}], "schema_path": ["discovery_consul", "access_info", "http_basic_auth_info", "passwd_url", "blindfold_secret_info"], "syntax": "block", "type": "object"}, {"aliases": ["clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:resources:discovery:properties:discovery_consul:access_info:http_basic_auth_info:passwd_url:clear_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-discovery_consul--access_info--http_basic_auth_info--passwd_url--clear_secret_info--url", "enforcement": "provider-schema", "group": "discovery_consul.access_info.http_basic_auth_info.passwd_url.clear_secret_info:RequiredObjectAttributes:url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:discovery_consul:access_info:http_basic_auth_info:passwd_url:clear_secret_info", "type": "requires"}], "schema_path": ["discovery_consul", "access_info", "http_basic_auth_info", "passwd_url", "clear_secret_info"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/discovery/properties/discovery_consul/access_info/http_basic_auth_info/passwd_url/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["discoveryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# discovery_consul.access_info.http_basic_auth_info.passwd_url

Breadcrumbs:

- [xcsh_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/)
- [discovery_consul](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/)
- [discovery_consul.access_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/access_info/)
- [discovery_consul.access_info.http_basic_auth_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/access_info/http_basic_auth_info/)
- discovery_consul.access_info.http_basic_auth_info.passwd_url

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
passwd_url {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/access_info/http_basic_auth_info/passwd_url/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/access_info/http_basic_auth_info/passwd_url/clear_secret_info/): complete subsection reference.

## Next pages

- [discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/access_info/http_basic_auth_info/passwd_url/blindfold_secret_info/)
- [discovery_consul.access_info.http_basic_auth_info.passwd_url.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/access_info/http_basic_auth_info/passwd_url/clear_secret_info/)
- [discovery_consul.access_info.http_basic_auth_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/access_info/http_basic_auth_info/)
- [xcsh_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/)
