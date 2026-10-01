---
page_title: "virtual_server.source_port"
subcategory: ""
description: "virtual_server.source_port for xcsh_application_profiles."
xcsh_docs: {"aliases": [], "body_bytes": 3038, "body_sha256": "sha256:1bd976b1f371471ff26613b826fbd0208e60b6f78784afe074293ed0edae5453", "child_ids": ["xcsh-docs:resources:application_profiles:properties:virtual_server:source_port:source_port_change", "xcsh-docs:resources:application_profiles:properties:virtual_server:source_port:source_port_preserve", "xcsh-docs:resources:application_profiles:properties:virtual_server:source_port:source_port_preserve_strict"], "collection_id": "xcsh-docs:resources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:resources:application_profiles:properties:virtual_server:source_port", "parent_id": "xcsh-docs:resources:application_profiles:properties:virtual_server", "path": "documentation/resources/application_profiles/properties/virtual_server/source_port/index.md", "provider_name": "application_profiles", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["virtual_server", "source_port"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/application_profiles/properties/virtual_server/source_port/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "virtual_server.source_port for xcsh_application_profiles.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["application_profilesCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# virtual_server.source_port

Breadcrumbs:

- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/)
- [virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/)
- virtual_server.source_port

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Specifies whether the system preserves the source port of the connection. The default is Preserve.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("source_port_change",
    "source_port_preserve"),
  validators.ConflictingObjectAttributes("source_port_change",
    "source_port_preserve_strict"),
  validators.ConflictingObjectAttributes("source_port_preserve",
    "source_port_preserve_strict")}
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
  "x-ves-oneof-field-source_port_choice": "[\"source_port_change\",\"source_port_preserve\",\"source_port_preserve_strict\"]"
}
```

Terraform syntax:

```terraform
source_port {
  # Configure direct properties listed below.
}
```

## Direct properties

- [source_port_change](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/source_port/source_port_change/): complete subsection reference.

- [source_port_preserve](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/source_port/source_port_preserve/): complete subsection reference.

- [source_port_preserve_strict](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/source_port/source_port_preserve_strict/): complete subsection reference.

## Next pages

- [virtual_server.source_port.source_port_change](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/source_port/source_port_change/)
- [virtual_server.source_port.source_port_preserve](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/source_port/source_port_preserve/)
- [virtual_server.source_port.source_port_preserve_strict](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/source_port/source_port_preserve_strict/)
- [virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/)
- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/)
