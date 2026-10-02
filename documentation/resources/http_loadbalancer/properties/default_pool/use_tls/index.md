---
page_title: "default_pool.use_tls"
subcategory: "Load Balancing"
description: "Upstream TLS Parameters."
xcsh_docs: {"aliases": ["default pool use tls"], "body_bytes": 9309, "body_sha256": "sha256:072f23cdd1226805c486fd233471af2a303af8f63c8383fa9321c22459791d00", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:default_session_key_caching", "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:disable_session_key_caching", "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:disable_sni", "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:no_mtls", "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:skip_server_verification", "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:tls_config", "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:use_host_header_as_sni", "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:use_mtls", "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:use_mtls_obj", "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:use_server_verification", "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:volterra_trusted_ca"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool", "path": "documentation/resources/http_loadbalancer/properties/default_pool/use_tls/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-017.md", "relationships": [{"anchor": "schema-default_pool--use_tls--max_session_keys", "enforcement": "provider-schema", "group": "default_pool.use_tls:ConflictingObjectAttributes:default_session_key_caching,max_session_keys", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls", "type": "conflicts"}, {"anchor": "schema-default_pool--use_tls--max_session_keys", "enforcement": "provider-schema", "group": "default_pool.use_tls:ConflictingObjectAttributes:disable_session_key_caching,max_session_keys", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls", "type": "conflicts"}, {"anchor": "schema-default_pool--use_tls--sni", "enforcement": "provider-schema", "group": "default_pool.use_tls:ConflictingObjectAttributes:disable_sni,sni", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls", "type": "conflicts"}, {"anchor": "schema-default_pool--use_tls--sni", "enforcement": "provider-schema", "group": "default_pool.use_tls:ConflictingObjectAttributes:sni,use_host_header_as_sni", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.use_tls:ConflictingObjectAttributes:default_session_key_caching,disable_session_key_caching", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:default_session_key_caching", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.use_tls:ConflictingObjectAttributes:default_session_key_caching,max_session_keys", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:default_session_key_caching", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.use_tls:ConflictingObjectAttributes:default_session_key_caching,disable_session_key_caching", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:disable_session_key_caching", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.use_tls:ConflictingObjectAttributes:disable_session_key_caching,max_session_keys", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:disable_session_key_caching", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.use_tls:ConflictingObjectAttributes:disable_sni,sni", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:disable_sni", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.use_tls:ConflictingObjectAttributes:disable_sni,use_host_header_as_sni", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:disable_sni", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.use_tls:ConflictingObjectAttributes:no_mtls,use_mtls", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:no_mtls", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.use_tls:ConflictingObjectAttributes:no_mtls,use_mtls_obj", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:no_mtls", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.use_tls:ConflictingObjectAttributes:skip_server_verification,use_server_verification", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:skip_server_verification", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.use_tls:ConflictingObjectAttributes:skip_server_verification,volterra_trusted_ca", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:skip_server_verification", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.use_tls:ConflictingObjectAttributes:disable_sni,use_host_header_as_sni", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:use_host_header_as_sni", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.use_tls:ConflictingObjectAttributes:sni,use_host_header_as_sni", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:use_host_header_as_sni", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.use_tls:ConflictingObjectAttributes:no_mtls,use_mtls", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:use_mtls", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.use_tls:ConflictingObjectAttributes:use_mtls,use_mtls_obj", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:use_mtls", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.use_tls:ConflictingObjectAttributes:no_mtls,use_mtls_obj", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:use_mtls_obj", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.use_tls:ConflictingObjectAttributes:use_mtls,use_mtls_obj", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:use_mtls_obj", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.use_tls:ConflictingObjectAttributes:skip_server_verification,use_server_verification", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:use_server_verification", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.use_tls:ConflictingObjectAttributes:use_server_verification,volterra_trusted_ca", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:use_server_verification", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.use_tls:ConflictingObjectAttributes:skip_server_verification,volterra_trusted_ca", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:volterra_trusted_ca", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.use_tls:ConflictingObjectAttributes:use_server_verification,volterra_trusted_ca", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:volterra_trusted_ca", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["default_pool", "use_tls"], "schema_version": 1, "sections": [{"aliases": ["default session key caching"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:default_session_key_caching", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_pool", "use_tls", "default_session_key_caching"], "syntax": "attribute", "type": "object"}, {"aliases": ["disable session key caching"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:disable_session_key_caching", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_pool", "use_tls", "disable_session_key_caching"], "syntax": "attribute", "type": "object"}, {"aliases": ["disable sni"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:disable_sni", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_pool", "use_tls", "disable_sni"], "syntax": "attribute", "type": "object"}, {"aliases": ["max session keys"], "anchor": "schema-default_pool--use_tls--max_session_keys", "description": "Exclusive with Number of session keys that are cached.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_pool", "use_tls", "max_session_keys"], "syntax": "attribute", "type": "number"}, {"aliases": ["no mtls"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:no_mtls", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_pool", "use_tls", "no_mtls"], "syntax": "attribute", "type": "object"}, {"aliases": ["skip server verification"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:skip_server_verification", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_pool", "use_tls", "skip_server_verification"], "syntax": "attribute", "type": "object"}, {"aliases": ["sni"], "anchor": "schema-default_pool--use_tls--sni", "description": "Exclusive with SNI value to be used.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_pool", "use_tls", "sni"], "syntax": "attribute", "type": "string"}, {"aliases": ["tls config"], "anchor": "section", "description": "This defines various OPTIONS to configure TLS configuration parameters.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:tls_config", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.use_tls.tls_config:ConflictingObjectAttributes:custom_security,default_security", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:tls_config:custom_security", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.use_tls.tls_config:ConflictingObjectAttributes:custom_security,low_security", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:tls_config:custom_security", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.use_tls.tls_config:ConflictingObjectAttributes:custom_security,medium_security", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:tls_config:custom_security", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.use_tls.tls_config:ConflictingObjectAttributes:custom_security,default_security", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:tls_config:default_security", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.use_tls.tls_config:ConflictingObjectAttributes:default_security,low_security", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:tls_config:default_security", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.use_tls.tls_config:ConflictingObjectAttributes:default_security,medium_security", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:tls_config:default_security", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.use_tls.tls_config:ConflictingObjectAttributes:custom_security,low_security", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:tls_config:low_security", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.use_tls.tls_config:ConflictingObjectAttributes:default_security,low_security", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:tls_config:low_security", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.use_tls.tls_config:ConflictingObjectAttributes:low_security,medium_security", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:tls_config:low_security", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.use_tls.tls_config:ConflictingObjectAttributes:custom_security,medium_security", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:tls_config:medium_security", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.use_tls.tls_config:ConflictingObjectAttributes:default_security,medium_security", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:tls_config:medium_security", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.use_tls.tls_config:ConflictingObjectAttributes:low_security,medium_security", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:tls_config:medium_security", "type": "conflicts"}], "schema_path": ["default_pool", "use_tls", "tls_config"], "syntax": "block", "type": "object"}, {"aliases": ["use host header as sni"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:use_host_header_as_sni", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_pool", "use_tls", "use_host_header_as_sni"], "syntax": "attribute", "type": "object"}, {"aliases": ["use mtls"], "anchor": "section", "description": "MTLS Client Certificate.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:use_mtls", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.use_tls.use_mtls:RequiredObjectAttributes:tls_certificates", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:use_mtls:tls_certificates", "type": "requires"}], "schema_path": ["default_pool", "use_tls", "use_mtls"], "syntax": "block", "type": "object"}, {"aliases": ["use mtls obj"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:use_mtls_obj", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-default_pool--use_tls--use_mtls_obj--name", "enforcement": "provider-schema", "group": "default_pool.use_tls.use_mtls_obj:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:use_mtls_obj", "type": "requires"}], "schema_path": ["default_pool", "use_tls", "use_mtls_obj"], "syntax": "block", "type": "object"}, {"aliases": ["use server verification"], "anchor": "section", "description": "Upstream TLS Validation Context.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:use_server_verification", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-default_pool--use_tls--use_server_verification--trusted_ca_url", "enforcement": "provider-schema", "group": "default_pool.use_tls.use_server_verification:ConflictingObjectAttributes:trusted_ca,trusted_ca_url", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:use_server_verification", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.use_tls.use_server_verification:ConflictingObjectAttributes:trusted_ca,trusted_ca_url", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:use_server_verification:trusted_ca", "type": "conflicts"}], "schema_path": ["default_pool", "use_tls", "use_server_verification"], "syntax": "block", "type": "object"}, {"aliases": ["volterra trusted ca"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:use_tls:volterra_trusted_ca", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_pool", "use_tls", "volterra_trusted_ca"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/default_pool/use_tls/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Upstream TLS Parameters.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_pool.use_tls

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [default_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/)
- default_pool.use_tls

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

TLS Parameters for Origin Servers. Upstream TLS Parameters.

Upstream description:

Upstream TLS Parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_session_key_caching",
    "disable_session_key_caching"),
  validators.ConflictingObjectAttributes("default_session_key_caching",
    "max_session_keys"),
  validators.ConflictingObjectAttributes("disable_session_key_caching",
    "max_session_keys"),
  validators.ConflictingObjectAttributes("disable_sni",
    "sni"),
  validators.ConflictingObjectAttributes("disable_sni",
    "use_host_header_as_sni"),
  validators.ConflictingObjectAttributes("no_mtls",
    "use_mtls"),
  validators.ConflictingObjectAttributes("no_mtls",
    "use_mtls_obj"),
  validators.ConflictingObjectAttributes("skip_server_verification",
    "use_server_verification"),
  validators.ConflictingObjectAttributes("skip_server_verification",
    "volterra_trusted_ca"),
  validators.ConflictingObjectAttributes("sni",
    "use_host_header_as_sni"),
  validators.ConflictingObjectAttributes("use_mtls",
    "use_mtls_obj"),
  validators.ConflictingObjectAttributes("use_server_verification",
    "volterra_trusted_ca")}
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
  "x-ves-oneof-field-max_session_keys_type": "[\"default_session_key_caching\",\"disable_session_key_caching\",\"max_session_keys\"]",
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\",\"use_mtls_obj\"]",
  "x-ves-oneof-field-server_validation_choice": "[\"skip_server_verification\",\"use_server_verification\",\"volterra_trusted_ca\"]",
  "x-ves-oneof-field-sni_choice": "[\"disable_sni\",\"sni\",\"use_host_header_as_sni\"]"
}
```

Terraform syntax:

```terraform
use_tls {
  # Configure direct properties listed below.
}
```

## Direct properties

- [default_session_key_caching](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/use_tls/default_session_key_caching/): complete subsection reference.

- [disable_session_key_caching](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/use_tls/disable_session_key_caching/): complete subsection reference.

- [disable_sni](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/use_tls/disable_sni/): complete subsection reference.

<a id="schema-default_pool--use_tls--max_session_keys"></a>

### max_session_keys property

Type: `"number"`. Optional.

Exclusive with \[default\_session\_key\_caching disable\_session\_key\_caching\] Number of session
keys that are cached.

Upstream description:

Exclusive with \[default\_session\_key\_caching disable\_session\_key\_caching\]

Number of session keys that are cached.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(2, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 2
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2",
    "ves.io.schema.rules.uint32.lte": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2",
    "ves.io.schema.rules.uint32.lte": "64"
  }
}
```

- [no_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/use_tls/no_mtls/): complete subsection reference.

- [skip_server_verification](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/use_tls/skip_server_verification/): complete subsection reference.

<a id="schema-default_pool--use_tls--sni"></a>

### sni property

Type: `"string"`. Optional.

Exclusive with \[disable\_sni use\_host\_header\_as\_sni\] SNI value to be used.

Upstream description:

Exclusive with \[disable\_sni use\_host\_header\_as\_sni\] SNI value to be used.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [tls_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/use_tls/tls_config/): complete subsection reference.

- [use_host_header_as_sni](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/use_tls/use_host_header_as_sni/): complete subsection reference.

- [use_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/use_tls/use_mtls/): complete subsection reference.

- [use_mtls_obj](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/use_tls/use_mtls_obj/): complete subsection reference.

- [use_server_verification](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/use_tls/use_server_verification/): complete subsection reference.

- [volterra_trusted_ca](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/use_tls/volterra_trusted_ca/): complete subsection reference.

## Next pages

- [default_pool.use_tls.default_session_key_caching](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/use_tls/default_session_key_caching/)
- [default_pool.use_tls.disable_session_key_caching](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/use_tls/disable_session_key_caching/)
- [default_pool.use_tls.disable_sni](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/use_tls/disable_sni/)
- [default_pool.use_tls.no_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/use_tls/no_mtls/)
- [default_pool.use_tls.skip_server_verification](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/use_tls/skip_server_verification/)
- [default_pool.use_tls.tls_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/use_tls/tls_config/)
- [default_pool.use_tls.use_host_header_as_sni](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/use_tls/use_host_header_as_sni/)
- [default_pool.use_tls.use_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/use_tls/use_mtls/)
- [default_pool.use_tls.use_mtls_obj](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/use_tls/use_mtls_obj/)
- [default_pool.use_tls.use_server_verification](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/use_tls/use_server_verification/)
- [default_pool.use_tls.volterra_trusted_ca](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/use_tls/volterra_trusted_ca/)
- [default_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
