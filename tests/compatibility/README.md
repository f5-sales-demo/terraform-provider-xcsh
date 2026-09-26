# Provider compatibility fixture

`v9-provider-schema/provider-schema.json.gz` is a deterministic gzip of the unmodified output of
`terraform providers schema -json` after Terraform installed
`registry.terraform.io/f5-sales-demo/xcsh` version `9.5.2`. The accompanying
configuration and lock file pin that provenance. The schema SHA-256 is
`f7fd2967ea6799de9d1b60e0937e11d630093428c7f26efc41d9ffe95262333f`.

Before the provider release surface was restored, the compatibility check
failed against the unmodified v11.3.0 source with removed baseline type names
and nine configured-field differences. The representative CSD configuration
also failed validation because `xcsh_addon_service_activation_status`,
`xcsh_namespace`, and `xcsh_protected_domain` were unavailable.

The nine field differences are reviewed v11 SMSv2 replacements. They retain
the system-scoped, non-forced v11 action contract and are documented in
`v9-provider-schema/reviewed-replacements.json`; every other v9.5.2
configurable field must remain compatible.
