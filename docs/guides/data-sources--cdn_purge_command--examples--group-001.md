---
page_title: "xcsh_cdn_purge_command examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cdn_purge_command examples."
---

# xcsh_cdn_purge_command examples

<a id="canonical-2122121333200302-2331113010221332-2131222102331332-1212022030030013-0300301001103103-1111302312103002-1313003100101122-0210003320133232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_cdn_purge_command](../data-sources/cdn_purge_command.md#canonical-1231330310300330-1311110203001010-3201101112230313-3011303001322021-2233001121001013-3002302102130103-3111000011322002-2220332233122201)
- Examples

<a id="canonical-0120033202101123-0012322322200030-2030332202311203-2021333232002002-3020223000012220-1222002023300113-0232220123101003-2031013002012112"></a>

### Complete configurations for `xcsh_cdn_purge_command`

- [Data source](data-sources--cdn_purge_command--examples--group-001.md#canonical-2211013302231123-0023210003112102-1111023330301123-0320113021003123-2323130002101121-2233213301221312-1231320333301230-0031221111310032): valid configuration.

<a id="canonical-2211013302231123-0023210003112102-1111023330301123-0320113021003123-2323130002101121-2233213301221312-1231320333301230-0031221111310032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_cdn_purge_command](../data-sources/cdn_purge_command.md#canonical-1231330310300330-1311110203001010-3201101112230313-3011303001322021-2233001121001013-3002302102130103-3111000011322002-2220332233122201)
- [Examples](data-sources--cdn_purge_command--examples--group-001.md#canonical-2122121333200302-2331113010221332-2131222102331332-1212022030030013-0300301001103103-1111302312103002-1313003100101122-0210003320133232)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_cdn_purge_command/data-source.tf`; digest `sha256:25cb8095a143a10866ff868fcae4396aaa9eca15b734c77fdd0d7054bddadd8a`.

```terraform
# CDNPurgeCommand Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing CDNPurgeCommand by name
data "xcsh_cdn_purge_command" "example" {
  name      = "example-cdn-purge-command"
  namespace = "staging"
}

output "cdn_purge_command_id" {
  value = data.xcsh_cdn_purge_command.example.id
}
```
