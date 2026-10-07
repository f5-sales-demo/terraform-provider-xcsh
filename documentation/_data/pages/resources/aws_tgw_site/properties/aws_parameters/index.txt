---
page_title: "aws_parameters"
subcategory: ""
description: "Setup AWS services VPC, transit gateway and site."
xcsh_docs: {"aliases": ["aws parameters"], "body_bytes": 12244, "body_sha256": "sha256:3690429e7124371def11a29a94700a66ce89329c372610e9404fd6dae6131243", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:admin_password", "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:aws_cred", "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:az_nodes", "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:custom_security_group", "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:disable_encryption", "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:disable_internet_vip", "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:enable_encryption", "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:enable_internet_vip", "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:existing_tgw", "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:f5xc_security_group", "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:new_tgw", "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:new_vpc", "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:no_worker_nodes", "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:reserved_tgw_cidr", "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:tgw_cidr"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters", "parent_id": "xcsh-docs:resources:aws_tgw_site:reference", "path": "documentation/resources/aws_tgw_site/properties/aws_parameters/index.md", "product": "distributed-cloud", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-0211211321130333-3310103111003032-2331313101231121-1000033231012332-1331312012302111-2003331003101111-3231100013121220-1221232021331020", "registry_path": "docs/guides/resources--aws_tgw_site--reference--group-001.md", "relationships": [{"anchor": "schema-aws_parameters--nodes_per_az", "enforcement": "provider-schema", "group": "aws_parameters:ConflictingObjectAttributes:no_worker_nodes,nodes_per_az", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters", "type": "conflicts"}, {"anchor": "schema-aws_parameters--nodes_per_az", "enforcement": "provider-schema", "group": "aws_parameters:ConflictingObjectAttributes:nodes_per_az,total_nodes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters", "type": "conflicts"}, {"anchor": "schema-aws_parameters--total_nodes", "enforcement": "provider-schema", "group": "aws_parameters:ConflictingObjectAttributes:no_worker_nodes,total_nodes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters", "type": "conflicts"}, {"anchor": "schema-aws_parameters--total_nodes", "enforcement": "provider-schema", "group": "aws_parameters:ConflictingObjectAttributes:nodes_per_az,total_nodes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters", "type": "conflicts"}, {"anchor": "schema-aws_parameters--vpc_id", "enforcement": "provider-schema", "group": "aws_parameters:ConflictingObjectAttributes:new_vpc,vpc_id", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_parameters:ConflictingObjectAttributes:custom_security_group,f5xc_security_group", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:custom_security_group", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_parameters:ConflictingObjectAttributes:disable_encryption,enable_encryption", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:disable_encryption", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_parameters:ConflictingObjectAttributes:disable_internet_vip,enable_internet_vip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:disable_internet_vip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_parameters:ConflictingObjectAttributes:disable_encryption,enable_encryption", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:enable_encryption", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_parameters:ConflictingObjectAttributes:disable_internet_vip,enable_internet_vip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:enable_internet_vip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_parameters:ConflictingObjectAttributes:existing_tgw,new_tgw", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:existing_tgw", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_parameters:ConflictingObjectAttributes:custom_security_group,f5xc_security_group", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:f5xc_security_group", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_parameters:ConflictingObjectAttributes:existing_tgw,new_tgw", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:new_tgw", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_parameters:ConflictingObjectAttributes:new_vpc,vpc_id", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:new_vpc", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_parameters:ConflictingObjectAttributes:no_worker_nodes,nodes_per_az", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:no_worker_nodes", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_parameters:ConflictingObjectAttributes:no_worker_nodes,total_nodes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:no_worker_nodes", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_parameters:ConflictingObjectAttributes:reserved_tgw_cidr,tgw_cidr", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:reserved_tgw_cidr", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_parameters:ConflictingObjectAttributes:reserved_tgw_cidr,tgw_cidr", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:tgw_cidr", "type": "conflicts"}, {"anchor": "schema-aws_parameters--aws_region", "enforcement": "provider-schema", "group": "aws_parameters:RequiredObjectAttributes:aws_region,az_nodes,instance_type,ssh_key", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters", "type": "requires"}, {"anchor": "schema-aws_parameters--instance_type", "enforcement": "provider-schema", "group": "aws_parameters:RequiredObjectAttributes:aws_region,az_nodes,instance_type,ssh_key", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters", "type": "requires"}, {"anchor": "schema-aws_parameters--ssh_key", "enforcement": "provider-schema", "group": "aws_parameters:RequiredObjectAttributes:aws_region,az_nodes,instance_type,ssh_key", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters", "type": "requires"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_parameters:RequiredObjectAttributes:aws_region,az_nodes,instance_type,ssh_key", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:az_nodes", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["aws_parameters"], "schema_version": 1, "sections": [{"aliases": ["aws parameters admin password"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:admin_password", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "aws_parameters.admin_password:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:admin_password:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_parameters.admin_password:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:admin_password:clear_secret_info", "type": "conflicts"}], "schema_path": ["aws_parameters", "admin_password"], "syntax": "block", "type": "object"}, {"aliases": ["aws parameters aws cred"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:aws_cred", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-aws_parameters--aws_cred--name", "enforcement": "provider-schema", "group": "aws_parameters.aws_cred:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:aws_cred", "type": "requires"}], "schema_path": ["aws_parameters", "aws_cred"], "syntax": "block", "type": "object"}, {"aliases": ["aws parameters aws region"], "anchor": "schema-aws_parameters--aws_region", "description": "AWS Region of your services VPC, where F5XC site will be deployed.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_parameters", "aws_region"], "syntax": "attribute", "type": "string"}, {"aliases": ["aws parameters az nodes"], "anchor": "section", "description": "Only Single AZ or Three AZ(s) nodes are supported currently.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:az_nodes", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "aws_parameters.az_nodes:ConflictingListObjectAttributes:inside_subnet,reserved_inside_subnet", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:az_nodes:inside_subnet", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_parameters.az_nodes:ConflictingListObjectAttributes:inside_subnet,reserved_inside_subnet", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:az_nodes:reserved_inside_subnet", "type": "conflicts"}, {"anchor": "schema-aws_parameters--az_nodes--aws_az_name", "enforcement": "provider-schema", "group": "aws_parameters.az_nodes:RequiredListObjectAttributes:aws_az_name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:az_nodes", "type": "requires"}], "schema_path": ["aws_parameters", "az_nodes"], "syntax": "block", "type": "object"}, {"aliases": ["aws parameters custom security group"], "anchor": "section", "description": "Enter pre created security groups for slo(Site Local Outside) and sli(Site Local Inside) interface. Supported only for sites deployed on existing VPC.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:custom_security_group", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["aws_parameters", "custom_security_group"], "syntax": "block", "type": "object"}, {"aliases": ["aws parameters disable encryption"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:disable_encryption", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_parameters", "disable_encryption"], "syntax": "attribute", "type": "object"}, {"aliases": ["aws parameters disable internet vip"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:disable_internet_vip", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_parameters", "disable_internet_vip"], "syntax": "attribute", "type": "object"}, {"aliases": ["aws parameters disk size"], "anchor": "schema-aws_parameters--disk_size", "description": "Node disk size for all node in the F5XC site. Unit is GiB.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_parameters", "disk_size"], "syntax": "attribute", "type": "number"}, {"aliases": ["aws parameters enable encryption"], "anchor": "section", "description": "Information related to disk encryption.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:enable_encryption", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-aws_parameters--enable_encryption--kms_key_id", "enforcement": "provider-schema", "group": "aws_parameters.enable_encryption:RequiredObjectAttributes:kms_key_id", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:enable_encryption", "type": "requires"}], "schema_path": ["aws_parameters", "enable_encryption"], "syntax": "block", "type": "object"}, {"aliases": ["aws parameters enable internet vip"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:enable_internet_vip", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_parameters", "enable_internet_vip"], "syntax": "attribute", "type": "object"}, {"aliases": ["aws parameters existing tgw"], "anchor": "section", "description": "Information needed for existing TGW.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:existing_tgw", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["aws_parameters", "existing_tgw"], "syntax": "block", "type": "object"}, {"aliases": ["aws parameters f5xc security group"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:f5xc_security_group", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_parameters", "f5xc_security_group"], "syntax": "attribute", "type": "object"}, {"aliases": ["aws parameters instance type"], "anchor": "schema-aws_parameters--instance_type", "description": "Instance size based on the performance.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_parameters", "instance_type"], "syntax": "attribute", "type": "string"}, {"aliases": ["aws parameters new tgw"], "anchor": "section", "description": "TGWParamsType.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:new_tgw", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "aws_parameters.new_tgw:ConflictingObjectAttributes:system_generated,user_assigned", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:new_tgw:system_generated", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_parameters.new_tgw:ConflictingObjectAttributes:system_generated,user_assigned", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:new_tgw:user_assigned", "type": "conflicts"}], "schema_path": ["aws_parameters", "new_tgw"], "syntax": "block", "type": "object"}, {"aliases": ["aws parameters new vpc"], "anchor": "section", "description": "Parameters to create new AWS VPC.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:new_vpc", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-aws_parameters--new_vpc--name_tag", "enforcement": "provider-schema", "group": "aws_parameters.new_vpc:ConflictingObjectAttributes:autogenerate,name_tag", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:new_vpc", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_parameters.new_vpc:ConflictingObjectAttributes:autogenerate,name_tag", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:new_vpc:autogenerate", "type": "conflicts"}, {"anchor": "schema-aws_parameters--new_vpc--primary_ipv4", "enforcement": "provider-schema", "group": "aws_parameters.new_vpc:RequiredObjectAttributes:primary_ipv4", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:new_vpc", "type": "requires"}], "schema_path": ["aws_parameters", "new_vpc"], "syntax": "block", "type": "object"}, {"aliases": ["aws parameters no worker nodes"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:no_worker_nodes", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_parameters", "no_worker_nodes"], "syntax": "attribute", "type": "object"}, {"aliases": ["aws parameters nodes per az"], "anchor": "schema-aws_parameters--nodes_per_az", "description": "Exclusive with Desired Worker Nodes Per AZ. Max limit is up to 21.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_parameters", "nodes_per_az"], "syntax": "attribute", "type": "number"}, {"aliases": ["aws parameters reserved tgw cidr"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:reserved_tgw_cidr", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_parameters", "reserved_tgw_cidr"], "syntax": "attribute", "type": "object"}, {"aliases": ["aws parameters ssh key"], "anchor": "schema-aws_parameters--ssh_key", "description": "Public SSH key for accessing nodes of the site.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_parameters", "ssh_key"], "syntax": "attribute", "type": "string"}, {"aliases": ["aws parameters tgw cidr"], "anchor": "section", "description": "Parameters for creating a new cloud subnet.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:tgw_cidr", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-aws_parameters--tgw_cidr--ipv4", "enforcement": "provider-schema", "group": "aws_parameters.tgw_cidr:RequiredObjectAttributes:ipv4", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:tgw_cidr", "type": "requires"}], "schema_path": ["aws_parameters", "tgw_cidr"], "syntax": "block", "type": "object"}, {"aliases": ["aws parameters total nodes"], "anchor": "schema-aws_parameters--total_nodes", "description": "Exclusive with Total number of worker nodes to be deployed across all AZ's used in the Site.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_parameters", "total_nodes"], "syntax": "attribute", "type": "number"}, {"aliases": ["aws parameters vpc id"], "anchor": "schema-aws_parameters--vpc_id", "description": "Exclusive with Existing VPC ID.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_parameters", "vpc_id"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/aws_parameters/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Setup AWS services VPC, transit gateway and site.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws_parameters

Breadcrumbs:

- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/)
- aws_parameters

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Setup AWS services VPC, transit gateway and site.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("aws_region",
    "az_nodes",
    "instance_type",
    "ssh_key"),
  validators.ConflictingObjectAttributes("custom_security_group",
    "f5xc_security_group"),
  validators.ConflictingObjectAttributes("disable_encryption",
    "enable_encryption"),
  validators.ConflictingObjectAttributes("disable_internet_vip",
    "enable_internet_vip"),
  validators.ConflictingObjectAttributes("existing_tgw",
    "new_tgw"),
  validators.ConflictingObjectAttributes("new_vpc",
    "vpc_id"),
  validators.ConflictingObjectAttributes("no_worker_nodes",
    "nodes_per_az"),
  validators.ConflictingObjectAttributes("no_worker_nodes",
    "total_nodes"),
  validators.ConflictingObjectAttributes("nodes_per_az",
    "total_nodes"),
  validators.ConflictingObjectAttributes("reserved_tgw_cidr",
    "tgw_cidr")}
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
  "x-ves-oneof-field-deployment": "[\"aws_cred\"]",
  "x-ves-oneof-field-encryption_choice": "[\"disable_encryption\",\"enable_encryption\"]",
  "x-ves-oneof-field-internet_vip_choice": "[\"disable_internet_vip\",\"enable_internet_vip\"]",
  "x-ves-oneof-field-security_group_choice": "[\"custom_security_group\",\"f5xc_security_group\"]",
  "x-ves-oneof-field-service_vpc_choice": "[\"new_vpc\",\"vpc_id\"]",
  "x-ves-oneof-field-tgw_choice": "[\"existing_tgw\",\"new_tgw\"]",
  "x-ves-oneof-field-tgw_cidr_choice": "[\"reserved_tgw_cidr\",\"tgw_cidr\"]",
  "x-ves-oneof-field-worker_nodes": "[\"no_worker_nodes\",\"nodes_per_az\",\"total_nodes\"]"
}
```

Terraform syntax:

```terraform
aws_parameters {
  # Configure direct properties listed below.
}
```

## Direct properties

- [admin_password](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/admin_password/): complete subsection reference.

- [aws_cred](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/aws_cred/): complete subsection reference.

<a id="schema-aws_parameters--aws_region"></a>

### aws_region property

Type: `"string"`. Optional.

AWS Region of your services VPC, where F5XC site will be deployed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

- [az_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/az_nodes/): complete subsection reference.

- [custom_security_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/custom_security_group/): complete subsection reference.

- [disable_encryption](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/disable_encryption/): complete subsection reference.

- [disable_internet_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/disable_internet_vip/): complete subsection reference.

<a id="schema-aws_parameters--disk_size"></a>

### disk_size property

Type: `"number"`. Optional.

Node disk size for all node in the F5XC site. Unit is GiB.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.AtMost(64000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 64000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "64000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "64000"
  }
}
```

- [enable_encryption](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/enable_encryption/): complete subsection reference.

- [enable_internet_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/enable_internet_vip/): complete subsection reference.

- [existing_tgw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/existing_tgw/): complete subsection reference.

- [f5xc_security_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/f5xc_security_group/): complete subsection reference.

<a id="schema-aws_parameters--instance_type"></a>

### instance_type property

Type: `"string"`. Optional.

AWS Instance Type for Node. Instance size based on the performance.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

- [new_tgw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/new_tgw/): complete subsection reference.

- [new_vpc](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/new_vpc/): complete subsection reference.

- [no_worker_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/no_worker_nodes/): complete subsection reference.

<a id="schema-aws_parameters--nodes_per_az"></a>

### nodes_per_az property

Type: `"number"`. Optional.

Exclusive with \[no\_worker\_nodes total\_nodes\] Desired Worker Nodes Per AZ. Max limit is up to
21.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 21),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 21,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "21"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "21"
  }
}
```

- [reserved_tgw_cidr](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/reserved_tgw_cidr/): complete subsection reference.

<a id="schema-aws_parameters--ssh_key"></a>

### ssh_key property

Type: `"string"`. Optional.

Public SSH key for accessing nodes of the site.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "8192",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "8192",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [tgw_cidr](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/tgw_cidr/): complete subsection reference.

<a id="schema-aws_parameters--total_nodes"></a>

### total_nodes property

Type: `"number"`. Optional.

Exclusive with \[no\_worker\_nodes nodes\_per\_az\] Total number of worker nodes to be deployed
across all AZ's used in the Site.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 61),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 61,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "61"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "61"
  }
}
```

<a id="schema-aws_parameters--vpc_id"></a>

### vpc_id property

Type: `"string"`. Optional.

Exclusive with \[new\_vpc\] Existing VPC ID.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```
