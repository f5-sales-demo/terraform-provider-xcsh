---
page_title: "palo_alto_fw_service"
subcategory: ""
description: "Palo Alto Networks VM-Series next-generation firewall configuration."
xcsh_docs: {"aliases": ["palo alto fw service"], "body_bytes": 16441, "body_sha256": "sha256:919a4167a909adfdcf7baa6fbf67b2f949fc84e83af2ab8dc66dd825ec15ee8c", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:auto_setup", "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:aws_tgw_site", "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:disable_panaroma", "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:pan_ami_bundle1", "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:pan_ami_bundle2", "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:panorama_server", "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:service_nodes"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service", "parent_id": "xcsh-docs:resources:nfv_service:reference", "path": "documentation/resources/nfv_service/properties/palo_alto_fw_service/index.md", "product": "distributed-cloud", "provider_name": "nfv_service", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1220032300320010-0322310310320112-3102020032000031-0230002231133113-1321103031132221-2112211003001211-2001103020013300-3102213200033033", "registry_path": "docs/guides/resources--nfv_service--reference--group-003.md", "relationships": [{"anchor": "schema-palo_alto_fw_service--ssh_key", "enforcement": "provider-schema", "group": "palo_alto_fw_service:ConflictingObjectAttributes:auto_setup,ssh_key", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "palo_alto_fw_service:ConflictingObjectAttributes:auto_setup,ssh_key", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:auto_setup", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "palo_alto_fw_service:ConflictingObjectAttributes:disable_panaroma,panorama_server", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:disable_panaroma", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "palo_alto_fw_service:ConflictingObjectAttributes:pan_ami_bundle1,pan_ami_bundle2", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:pan_ami_bundle1", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "palo_alto_fw_service:ConflictingObjectAttributes:pan_ami_bundle1,pan_ami_bundle2", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:pan_ami_bundle2", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "palo_alto_fw_service:ConflictingObjectAttributes:disable_panaroma,panorama_server", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:panorama_server", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["palo_alto_fw_service"], "schema_version": 1, "sections": [{"aliases": ["auto setup"], "anchor": "section", "description": "For auto-setup, SSH public and pvt keys are needed. Using the given config user, SSH and API access will be configured.", "document_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:auto_setup", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-palo_alto_fw_service--auto_setup--admin_username", "enforcement": "provider-schema", "group": "palo_alto_fw_service.auto_setup:RequiredObjectAttributes:admin_username", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:auto_setup", "type": "requires"}], "schema_path": ["palo_alto_fw_service", "auto_setup"], "syntax": "block", "type": "object"}, {"aliases": ["aws tgw site"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:aws_tgw_site", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-palo_alto_fw_service--aws_tgw_site--name", "enforcement": "provider-schema", "group": "palo_alto_fw_service.aws_tgw_site:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:aws_tgw_site", "type": "requires"}], "schema_path": ["palo_alto_fw_service", "aws_tgw_site"], "syntax": "block", "type": "object"}, {"aliases": ["disable panaroma"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:disable_panaroma", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["palo_alto_fw_service", "disable_panaroma"], "syntax": "attribute", "type": "object"}, {"aliases": ["instance type"], "anchor": "schema-palo_alto_fw_service--instance_type", "description": "- PALO_ALTO_FW_AWS_INSTANCE_TYPE_M4_XLARGE: m4.xlarge - PALO_ALTO_FW_AWS_INSTANCE_TYPE_M4_2XLARGE: m4.2xlarge - PALO_ALTO_FW_AWS_INSTANCE_TYPE_M4_4XLARGE: m4.4xlarge - PALO_ALTO_FW_AWS_INSTANCE_TYPE_M5_LARGE: m5.large - PALO_ALTO_FW_AWS_INSTANCE_TYPE_M5_XLARGE: m5.xlarge - PALO_ALTO_FW_AWS_INSTANCE_TYPE_M5_2XLARGE: m5.", "document_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["palo_alto_fw_service", "instance_type"], "syntax": "attribute", "type": "string"}, {"aliases": ["pan ami bundle1"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:pan_ami_bundle1", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["palo_alto_fw_service", "pan_ami_bundle1"], "syntax": "attribute", "type": "object"}, {"aliases": ["pan ami bundle2"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:pan_ami_bundle2", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["palo_alto_fw_service", "pan_ami_bundle2"], "syntax": "attribute", "type": "object"}, {"aliases": ["panorama server"], "anchor": "section", "description": "Panorama Server Type.", "document_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:panorama_server", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-palo_alto_fw_service--panorama_server--server", "enforcement": "provider-schema", "group": "palo_alto_fw_service.panorama_server:RequiredObjectAttributes:server", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:panorama_server", "type": "requires"}], "schema_path": ["palo_alto_fw_service", "panorama_server"], "syntax": "block", "type": "object"}, {"aliases": ["service nodes"], "anchor": "section", "description": "Configuration parameter for service nodes.", "document_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:service_nodes", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "palo_alto_fw_service.service_nodes:RequiredObjectAttributes:nodes", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service:service_nodes:nodes", "type": "requires"}], "schema_path": ["palo_alto_fw_service", "service_nodes"], "syntax": "block", "type": "object"}, {"aliases": ["ssh key"], "anchor": "schema-palo_alto_fw_service--ssh_key", "description": "Exclusive with Setup Authorized Public SSH key. User will be able to SSH to the vmseries nodes using its corresponding SSH private key.", "document_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["palo_alto_fw_service", "ssh_key"], "syntax": "attribute", "type": "string"}, {"aliases": ["tags"], "anchor": "schema-palo_alto_fw_service--tags", "description": "AWS Tags is a label consisting of a user-defined key and value. It helps to manage, identify, organize, search for, and filter resources in AWS console.", "document_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["palo_alto_fw_service", "tags"], "syntax": "attribute", "type": "map"}, {"aliases": ["version"], "anchor": "schema-palo_alto_fw_service--version", "description": "PAN-OS version.", "document_id": "xcsh-docs:resources:nfv_service:properties:palo_alto_fw_service", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["palo_alto_fw_service", "version"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nfv_service/properties/palo_alto_fw_service/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Palo Alto Networks VM-Series next-generation firewall configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# palo_alto_fw_service

Breadcrumbs:

- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/)
- palo_alto_fw_service

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Palo Alto Networks VM-Series next-generation firewall configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("auto_setup",
    "ssh_key"),
  validators.ConflictingObjectAttributes("disable_panaroma",
    "panorama_server"),
  validators.ConflictingObjectAttributes("pan_ami_bundle1",
    "pan_ami_bundle2")}
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
  "x-ves-oneof-field-ami_choice": "[\"pan_ami_bundle1\",\"pan_ami_bundle2\"]",
  "x-ves-oneof-field-panaroma_connection": "[\"disable_panaroma\",\"panorama_server\"]",
  "x-ves-oneof-field-setup_options": "[\"auto_setup\",\"ssh_key\"]"
}
```

Terraform syntax:

```terraform
palo_alto_fw_service {
  # Configure direct properties listed below.
}
```

## Direct properties

- [auto_setup](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/palo_alto_fw_service/auto_setup/): complete subsection reference.

- [aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/palo_alto_fw_service/aws_tgw_site/): complete subsection reference.

- [disable_panaroma](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/palo_alto_fw_service/disable_panaroma/): complete subsection reference.

<a id="schema-palo_alto_fw_service--instance_type"></a>

### instance_type property

Type: `"string"`. Optional.

\[Enum:
PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M4\_XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M4\_2XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M4\_4XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_LARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_2XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_4XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_12XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5N\_LARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5N\_XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5N\_2XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5N\_4XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C4\_LARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C4\_XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C4\_2XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C4\_4XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C4\_8XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_LARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_2XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_4XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_9XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_18XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_LARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_2XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_4XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_9XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_18XLARGE|PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_R5\_2XLARGE\]
&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M4\_XLARGE: m4.xlarge -
PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M4\_2XLARGE: m4.2xlarge -
PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M4\_4XLARGE: m4.4xlarge -
PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_LARGE: m5.large -
PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_XLARGE: m5.xlarge .. Possible values are
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M4\_XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M4\_2XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M4\_4XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_LARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_2XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_4XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_12XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5N\_LARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5N\_XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5N\_2XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5N\_4XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C4\_LARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C4\_XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C4\_2XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C4\_4XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C4\_8XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_LARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_2XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_4XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_9XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_18XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_LARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_2XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_4XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_9XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_18XLARGE\`,
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_R5\_2XLARGE\`. Defaults to
\`PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M4\_XLARGE\`.

Upstream description:

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M4\_XLARGE: m4.xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M4\_2XLARGE: m4.2xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M4\_4XLARGE: m4.4xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_LARGE: m5.large

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_XLARGE: m5.xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_2XLARGE: m5.2xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_4XLARGE: m5.4xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5\_12XLARGE: m5.12xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5N\_LARGE: m5n.large

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5N\_XLARGE: m5n.xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5N\_2XLARGE: m5n.2xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_M5N\_4XLARGE: m5n.4xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C4\_LARGE: c4.large

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C4\_XLARGE: c4.xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C4\_2XLARGE: c4.2xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C4\_4XLARGE: c4.4xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C4\_8XLARGE: c4.8xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_LARGE: c5.large

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_XLARGE: c5.xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_2XLARGE: c5.2xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_4XLARGE: c5.4xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_9XLARGE: c5.9xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5\_18XLARGE: c5.18xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_LARGE: c5n.large

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_XLARGE: c5n.xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_2XLARGE: c5n.2xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_4XLARGE: c5n.4xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_9XLARGE: c5n.9xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_C5N\_18XLARGE: c5n.18xlarge

&#8203;- PALO\_ALTO\_FW\_AWS\_INSTANCE\_TYPE\_R5\_2XLARGE: r5.2xlarge.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("PALO_ALTO_FW_AWS_INSTANCE_TYPE_M4_XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M4_2XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M4_4XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M5_LARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M5_XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M5_2XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M5_4XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M5_12XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M5N_LARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M5N_XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M5N_2XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M5N_4XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C4_LARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C4_XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C4_2XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C4_4XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C4_8XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5_LARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5_XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5_2XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5_4XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5_9XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5_18XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5N_LARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5N_XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5N_2XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5N_4XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5N_9XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5N_18XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_R5_2XLARGE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M4_XLARGE",
  "enum": [
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M4_XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M4_2XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M4_4XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M5_LARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M5_XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M5_2XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M5_4XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M5_12XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M5N_LARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M5N_XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M5N_2XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_M5N_4XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C4_LARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C4_XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C4_2XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C4_4XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C4_8XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5_LARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5_XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5_2XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5_4XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5_9XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5_18XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5N_LARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5N_XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5N_2XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5N_4XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5N_9XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_C5N_18XLARGE",
    "PALO_ALTO_FW_AWS_INSTANCE_TYPE_R5_2XLARGE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [pan_ami_bundle1](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/palo_alto_fw_service/pan_ami_bundle1/): complete subsection reference.

- [pan_ami_bundle2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/palo_alto_fw_service/pan_ami_bundle2/): complete subsection reference.

- [panorama_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/palo_alto_fw_service/panorama_server/): complete subsection reference.

- [service_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/palo_alto_fw_service/service_nodes/): complete subsection reference.

<a id="schema-palo_alto_fw_service--ssh_key"></a>

### ssh_key property

Type: `"string"`. Optional.

Exclusive with \[auto\_setup\] Setup Authorized Public SSH key. User will be able to SSH to the
vmseries nodes using its corresponding SSH private key.

Upstream description:

Exclusive with \[auto\_setup\] Setup Authorized Public SSH key. User will be able to SSH to the
vmseries nodes using its corresponding SSH private key.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 8192),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8192,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8192,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8192",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8192",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="schema-palo_alto_fw_service--tags"></a>

### tags property

Type: `["map", "string"]`. Optional.

AWS Tags is a label consisting of a user-defined key and value. It helps to manage, identify,
organize, search for, and filter resources in AWS console.

Upstream description:

AWS Tags is a label consisting of a user-defined key and value. It helps to manage, identify,
organize, search for, and filter resources in AWS console.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "127",
    "ves.io.schema.rules.map.max_pairs": "40",
    "ves.io.schema.rules.map.values.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "127",
    "ves.io.schema.rules.map.max_pairs": "40",
    "ves.io.schema.rules.map.values.string.max_len": "255"
  }
}
```

<a id="schema-palo_alto_fw_service--version"></a>

### version property

Type: `"string"`. Optional.

\[Enum: 11.0.0\] PAN VM-Series version. PAN-OS version. The only possible value is \`11.0.0\`.

Upstream description:

PAN-OS version.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("11.0.0"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "11.0.0"
  ],
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"11.0.0\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"11.0.0\\\"]"
  }
}
```

## Next pages

- [palo_alto_fw_service.auto_setup](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/palo_alto_fw_service/auto_setup/)
- [palo_alto_fw_service.aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/palo_alto_fw_service/aws_tgw_site/)
- [palo_alto_fw_service.disable_panaroma](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/palo_alto_fw_service/disable_panaroma/)
- [palo_alto_fw_service.pan_ami_bundle1](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/palo_alto_fw_service/pan_ami_bundle1/)
- [palo_alto_fw_service.pan_ami_bundle2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/palo_alto_fw_service/pan_ami_bundle2/)
- [palo_alto_fw_service.panorama_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/palo_alto_fw_service/panorama_server/)
- [palo_alto_fw_service.service_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/palo_alto_fw_service/service_nodes/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/)
- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/)
