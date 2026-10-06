---
page_title: "xcsh_origin_pool examples"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_origin_pool examples."
---

# xcsh_origin_pool examples

<a id="canonical-0000311122100023-3302033101223232-1212303232002022-2202323221002012-2312220110133231-0212131003330311-3211202222332020-3322333102120030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- Examples

<a id="canonical-1231103230313310-3133132312031302-0230312123331312-0132131320122322-1230303312020220-2102013211203301-1010110133203203-2310020323012101"></a>

### Complete configurations for `xcsh_origin_pool`

- [Labels update](resources--origin_pool--examples--group-001.md#canonical-1121331123332011-0030013012030301-0323033030212313-1123120031123301-2213300211123332-1312212023121130-1033012130222210-3320011010230103): valid configuration.

- [Multiple origins](resources--origin_pool--examples--group-001.md#canonical-3233022010022103-0131333231223300-2112303032010131-2303302301330233-2232110202113013-3300113312232223-3100330113013101-1232322222220303): valid configuration.

- [Nested labels](resources--origin_pool--examples--group-001.md#canonical-2012132311320030-0222302331123230-3111030023131032-0022302033301311-3302201233221101-2031232202313322-1010131321002110-2000223021131332): valid configuration.

- [Port](resources--origin_pool--examples--group-001.md#canonical-2033033132020302-1000101012303013-0301322021333133-2131303012132211-2211011303031212-1233013200323322-0313223000323133-0212103003310103): valid configuration.

- [Public ip](resources--origin_pool--examples--group-001.md#canonical-0130130230222231-2122330321101001-3002302010012230-3310302321220300-3103121001032121-0223122131022201-2022032000120302-0001031222111212): valid configuration.

- [Resource](resources--origin_pool--examples--group-001.md#canonical-2030033210101022-2320311212030202-0001133120032331-1201133130223313-1320023002203021-0000300133122232-3112200112030310-3313211330330102): valid configuration.

- [With labels](resources--origin_pool--examples--group-001.md#canonical-2311121333023321-0221132131023221-1102302002223221-2320012302100233-0020022101223001-1023222202000033-2322203032102330-3002213122132033): valid configuration.

<a id="canonical-1121331123332011-0030013012030301-0323033030212313-1123120031123301-2213300211123332-1312212023121130-1033012130222210-3320011010230103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Labels update example

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Examples](resources--origin_pool--examples--group-001.md#canonical-0000311122100023-3302033101223232-1212303232002022-2202323221002012-2312220110133231-0212131003330311-3211202222332020-3322333102120030)
- Labels update

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_origin_pool/labels-update.tf`; digest `sha256:d2a6e4ca39384f01c4d80efec2f06fdf2e82a0293580f75d7c78c1f61e96c022`.

```terraform
# LabelsUpdate — Acceptance-test-derived Configuration
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

resource "xcsh_origin_pool" "test" {
  name      = "example"
  namespace = "system"

  port = 443

  labels = {
    environment = "example-value"
  }

  origin_servers {
    public_name {
      dns_name = "example.com"
    }
  }

  no_tls                = {}
  same_as_endpoint_port = {}
}
```

<a id="canonical-3233022010022103-0131333231223300-2112303032010131-2303302301330233-2232110202113013-3300113312232223-3100330113013101-1232322222220303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Multiple origins example

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Examples](resources--origin_pool--examples--group-001.md#canonical-0000311122100023-3302033101223232-1212303232002022-2202323221002012-2312220110133231-0212131003330311-3211202222332020-3322333102120030)
- Multiple origins

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_origin_pool/multiple-origins.tf`; digest `sha256:bd8a184b9c325d4b7310968c1937c2417c67cf0e0a8113c2bf868b7ccb8905ff`.

```terraform
# MultipleOrigins — Acceptance-test-derived Configuration
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

resource "xcsh_origin_pool" "test" {
  name      = "example"
  namespace = "system"

  port = 443

  origin_servers {
    public_name {
      dns_name = "backend1.example.com"
    }
  }

  origin_servers {
    public_name {
      dns_name = "backend2.example.com"
    }
  }

  no_tls                = {}
  same_as_endpoint_port = {}
}
```

<a id="canonical-2012132311320030-0222302331123230-3111030023131032-0022302033301311-3302201233221101-2031232202313322-1010131321002110-2000223021131332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Nested labels example

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Examples](resources--origin_pool--examples--group-001.md#canonical-0000311122100023-3302033101223232-1212303232002022-2202323221002012-2312220110133231-0212131003330311-3211202222332020-3322333102120030)
- Nested labels

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_origin_pool/nested-labels.tf`; digest `sha256:2a2c2b7a5d041bafee98f9e4a0464b2e4b5c967ac5f5b30fa2f474ff4f6ef535`.

```terraform
# NestedLabels — Acceptance-test-derived Configuration
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

resource "xcsh_origin_pool" "test" {
  name      = "example"
  namespace = "system"

  port = 8080

  origin_servers {
    public_ip {
      ip = "192.0.2.1"
    }
    labels = {
      "env" = "test"
      "app" = "demo"
    }
  }

  no_tls                = {}
  same_as_endpoint_port = {}
}
```

<a id="canonical-2033033132020302-1000101012303013-0301322021333133-2131303012132211-2211011303031212-1233013200323322-0313223000323133-0212103003310103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Port example

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Examples](resources--origin_pool--examples--group-001.md#canonical-0000311122100023-3302033101223232-1212303232002022-2202323221002012-2312220110133231-0212131003330311-3211202222332020-3322333102120030)
- Port

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_origin_pool/port.tf`; digest `sha256:3852ca8c70ab22d421232a89d1c20cf56857ad2339faf797aa538bea24c6a139`.

```terraform
# Port — Acceptance-test-derived Configuration
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

resource "xcsh_origin_pool" "test" {
  name      = "example"
  namespace = "system"

  port = 443

  origin_servers {
    public_name {
      dns_name = "example.com"
    }
  }

  no_tls                = {}
  same_as_endpoint_port = {}
}
```

<a id="canonical-0130130230222231-2122330321101001-3002302010012230-3310302321220300-3103121001032121-0223122131022201-2022032000120302-0001031222111212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Public ip example

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Examples](resources--origin_pool--examples--group-001.md#canonical-0000311122100023-3302033101223232-1212303232002022-2202323221002012-2312220110133231-0212131003330311-3211202222332020-3322333102120030)
- Public ip

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_origin_pool/public-ip.tf`; digest `sha256:1dade891c9c99d46213e707ba46387f7c7b5eed90053ec8616b5c928911c7d54`.

```terraform
# PublicIp — Acceptance-test-derived Configuration
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

resource "xcsh_origin_pool" "test" {
  name      = "example"
  namespace = "system"

  port = 8080

  origin_servers {
    public_ip {
      ip = "192.0.2.1"
    }
  }

  no_tls                = {}
  same_as_endpoint_port = {}
}
```

<a id="canonical-2030033210101022-2320311212030202-0001133120032331-1201133130223313-1320023002203021-0000300133122232-3112200112030310-3313211330330102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Examples](resources--origin_pool--examples--group-001.md#canonical-0000311122100023-3302033101223232-1212303232002022-2202323221002012-2312220110133231-0212131003330311-3211202222332020-3322333102120030)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_origin_pool/resource.tf`; digest `sha256:d9f0fe93931beb085130f6173ace667305af3adbe1e9448009e07911a7af8fa0`.

```terraform
# OriginPool Resource Example
# Manages a Origin Pool resource in F5 Distributed Cloud for defining backend server pools for load balancer targets.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic OriginPool configuration
resource "xcsh_origin_pool" "example" {
  name      = "example-origin-pool"
  namespace = "staging"
}
```

<a id="canonical-2311121333023321-0221132131023221-1102302002223221-2320012302100233-0020022101223001-1023222202000033-2322203032102330-3002213122132033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## With labels example

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Examples](resources--origin_pool--examples--group-001.md#canonical-0000311122100023-3302033101223232-1212303232002022-2202323221002012-2312220110133231-0212131003330311-3211202222332020-3322333102120030)
- With labels

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_origin_pool/with-labels.tf`; digest `sha256:492fffafa075a1f28ad253873263057c634ae541eb762e94f4e227737c234dcd`.

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

resource "xcsh_origin_pool" "test" {
  name        = "example"
  namespace   = "system"
  description = "Test origin pool"

  port = 443

  labels = {
    environment = "test"
    team        = "platform"
  }

  origin_servers {
    public_name {
      dns_name = "example.com"
    }
  }

  no_tls                = {}
  same_as_endpoint_port = {}
}
```
