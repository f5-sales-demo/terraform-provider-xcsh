---
page_title: "xcsh_service_policy examples"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_service_policy examples."
---

# xcsh_service_policy examples

<a id="canonical-1220321203332321-0223210020200002-3001003033013233-2101111202111230-3201121300222213-2000333221301112-3223311102130320-3311232312310321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- Examples

<a id="canonical-0200001220012110-0233213020033320-3222020030231001-2103303032300023-2201232130113113-2232232300111011-1332330202123311-3111021212231223"></a>

### Complete configurations for `xcsh_service_policy`

- [Allow list](resources--service_policy--examples--group-001.md#canonical-0233102322233123-0311301020032112-3131002122010220-2101203223123121-1232221122021101-0332213313021310-3122030132123230-0030312331210132): valid configuration.

- [Deny all](resources--service_policy--examples--group-001.md#canonical-1010220320132302-1032132231102012-0003122002201032-1230100323323112-3023322323032203-2100130100220132-2002320201331033-0211301312101222): valid configuration.

- [Deny list](resources--service_policy--examples--group-001.md#canonical-1020102021132323-1231031122030310-2211013123032313-2032332201010211-2203101100303003-3003023331132200-0000320211100120-2102211112223003): valid configuration.

- [Resource](resources--service_policy--examples--group-001.md#canonical-3320211212231301-0113220203302033-3022003101032102-0310113001330033-2333330111020131-2333202213330203-0101312123011230-3232302212131020): valid configuration.

- [With labels](resources--service_policy--examples--group-001.md#canonical-1100130121131320-3002000332012103-0130032200200313-1121102013220022-2000000010320213-0033103010320113-0313201323131001-2322131311220113): valid configuration.

<a id="canonical-0233102322233123-0311301020032112-3131002122010220-2101203223123121-1232221122021101-0332213313021310-3122030132123230-0030312331210132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Allow list example

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Examples](resources--service_policy--examples--group-001.md#canonical-1220321203332321-0223210020200002-3001003033013233-2101111202111230-3201121300222213-2000333221301112-3223311102130320-3311232312310321)
- Allow list

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_service_policy/allow-list.tf`; digest `sha256:e3e5a70e5d1b6c7d8436b9826fb90da18093ff0f1b51c42b00f77c75bb40e353`.

```terraform
# AllowList — Acceptance-test-derived Configuration
# Extracted from an acceptance test helper.
# No new live API validation is claimed.

terraform {
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

resource "xcsh_service_policy" "test" {
  name      = "example"
  namespace = "system"

  # Allow list with IP prefix
  allow_list {
    prefix_list {
      prefixes = ["10.0.0.0/8", "192.168.0.0/16"]
    }
    default_action_deny = {}
  }

  # Apply to any server
  any_server = {}
}
```

<a id="canonical-1010220320132302-1032132231102012-0003122002201032-1230100323323112-3023322323032203-2100130100220132-2002320201331033-0211301312101222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Deny all example

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Examples](resources--service_policy--examples--group-001.md#canonical-1220321203332321-0223210020200002-3001003033013233-2101111202111230-3201121300222213-2000333221301112-3223311102130320-3311232312310321)
- Deny all

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_service_policy/deny-all.tf`; digest `sha256:9da9688512df86906769c34d1e33802dc2db45bff21587219ced7d18b6172282`.

```terraform
# DenyAll — Acceptance-test-derived Configuration
# Extracted from an acceptance test helper.
# No new live API validation is claimed.

terraform {
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

resource "xcsh_service_policy" "test" {
  name      = "example"
  namespace = "system"

  # Deny all requests
  deny_all_requests = {}

  # Apply to any server
  any_server = {}
}
```

<a id="canonical-1020102021132323-1231031122030310-2211013123032313-2032332201010211-2203101100303003-3003023331132200-0000320211100120-2102211112223003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Deny list example

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Examples](resources--service_policy--examples--group-001.md#canonical-1220321203332321-0223210020200002-3001003033013233-2101111202111230-3201121300222213-2000333221301112-3223311102130320-3311232312310321)
- Deny list

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_service_policy/deny-list.tf`; digest `sha256:6000e015862d1bb25f8e6b5bc94664685fdab920f11f2608c4f47416c65ff393`.

```terraform
# DenyList — Acceptance-test-derived Configuration
# Extracted from an acceptance test helper.
# No new live API validation is claimed.

terraform {
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

resource "xcsh_service_policy" "test" {
  name      = "example"
  namespace = "system"

  deny_list {
    prefix_list {
      prefixes = ["172.16.0.0/12"]
    }
    default_action_allow = {}
  }

  any_server = {}
}
```

<a id="canonical-3320211212231301-0113220203302033-3022003101032102-0310113001330033-2333330111020131-2333202213330203-0101312123011230-3232302212131020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Examples](resources--service_policy--examples--group-001.md#canonical-1220321203332321-0223210020200002-3001003033013233-2101111202111230-3201121300222213-2000333221301112-3223311102130320-3311232312310321)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_service_policy/resource.tf`; digest `sha256:e0fbf1c5446df6211fe69296e8020455f7d5f52c63778a8e6218ac43e92a6906`.

```terraform
# ServicePolicy Resource Example
# Manages service_policy creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic ServicePolicy configuration
resource "xcsh_service_policy" "example" {
  name      = "example-service-policy"
  namespace = "staging"
}
```

<a id="canonical-1100130121131320-3002000332012103-0130032200200313-1121102013220022-2000000010320213-0033103010320113-0313201323131001-2322131311220113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## With labels example

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md#canonical-3130230113123200-0200122303013133-3003310003113113-2022033310201211-3303333001012130-2333023010121030-2132303102102003-2201111001022301)
- [Examples](resources--service_policy--examples--group-001.md#canonical-1220321203332321-0223210020200002-3001003033013233-2101111202111230-3201121300222213-2000333221301112-3223311102130320-3311232312310321)
- With labels

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_service_policy/with-labels.tf`; digest `sha256:27966c953ab74ec12398f4c688232398a984b67f1bec6a9179d66139d051b625`.

```terraform
# WithLabels — Acceptance-test-derived Configuration
# Extracted from an acceptance test helper.
# No new live API validation is claimed.

terraform {
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

resource "xcsh_service_policy" "test" {
  name        = "example"
  namespace   = "system"
  description = "Test service policy"

  labels = {
    environment = "test"
    team        = "security"
  }

  # Allow all requests
  allow_all_requests = {}

  # Apply to any server
  any_server = {}
}
```
