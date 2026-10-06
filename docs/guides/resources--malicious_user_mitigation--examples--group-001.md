---
page_title: "xcsh_malicious_user_mitigation examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_malicious_user_mitigation examples."
---

# xcsh_malicious_user_mitigation examples

<a id="canonical-0122333121121121-2133233012323031-2100322101123321-3101010121133222-1002313333123011-2132000233033313-0320000203211130-1132123110011023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md#canonical-1101131202021321-3221321112031101-3221121031010113-2211012203220000-0302111121132132-3121301122233133-1013203002033303-3311201201113123)
- Examples

<a id="canonical-2312332202203230-1201012300022001-3002202013200003-0303121323012020-2330103322313002-1223212213110201-1221222131013132-2123322022110220"></a>

### Complete configurations for `xcsh_malicious_user_mitigation`

- [All attributes](resources--malicious_user_mitigation--examples--group-001.md#canonical-1213311311203230-3220000000201113-3120133333033102-3320211002020021-0013323130120321-1002021332120123-0223222331233022-3133103001133121): valid configuration.

- [Block](resources--malicious_user_mitigation--examples--group-001.md#canonical-1021333202112011-1013203201111331-1231122032023311-3131300231110212-2231232021112203-2020100100110122-0320332111201312-2313030011321112): valid configuration.

- [Captcha](resources--malicious_user_mitigation--examples--group-001.md#canonical-0032002003213131-3101211313010231-1223010023132031-2211303333320030-2202213012323012-1301021111132001-1131030212123103-2100312220110012): valid configuration.

- [Js challenge](resources--malicious_user_mitigation--examples--group-001.md#canonical-3213332111110303-0233120032031331-3130132222023022-3233112232230001-1110123013010013-2020020120033103-1130312223101333-3203101321203012): valid configuration.

- [Resource](resources--malicious_user_mitigation--examples--group-001.md#canonical-3301001202201031-1213330213031310-1333230003311100-2013300321320202-0113311010311321-0233322120202133-3020120221033023-3023220011010101): valid configuration.

- [With annotations](resources--malicious_user_mitigation--examples--group-001.md#canonical-0332011323203221-2212021231330001-3022103131310300-1030022012220312-0103300311201021-3223103001302132-0300002300321220-3012322030320300): valid configuration.

- [With description](resources--malicious_user_mitigation--examples--group-001.md#canonical-0122001301220132-0333321133231321-2301323122121103-2200220230231102-1321331121211023-1120012130313133-2321011220121032-1300222031323120): valid configuration.

- [With labels](resources--malicious_user_mitigation--examples--group-001.md#canonical-3020323201221023-0013011020311132-0312111323302223-3333233023023120-0303012131030212-0213211312221323-1101020303111311-2212033020133012): valid configuration.

- [With mitigation type](resources--malicious_user_mitigation--examples--group-001.md#canonical-3202001211201031-1231330122022303-3311223013133003-2000101313231321-2130030101220033-0132303322101130-2333223130131021-1111222001302220): valid configuration.

<a id="canonical-1213311311203230-3220000000201113-3120133333033102-3320211002020021-0013323130120321-1002021332120123-0223222331233022-3133103001133121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## All attributes example

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md#canonical-1101131202021321-3221321112031101-3221121031010113-2211012203220000-0302111121132132-3121301122233133-1013203002033303-3311201201113123)
- [Examples](resources--malicious_user_mitigation--examples--group-001.md#canonical-0122333121121121-2133233012323031-2100322101123321-3101010121133222-1002313333123011-2132000233033313-0320000203211130-1132123110011023)
- All attributes

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_malicious_user_mitigation/all-attributes.tf`; digest `sha256:e1a870d0ac8f57cd049b19713024efa0b33bb0d6f7d3e2cc75464e43a5026739`.

```terraform
# AllAttributes — Acceptance-test-derived Configuration
# Extracted from an acceptance test helper.
# No new live API validation is claimed.

terraform {
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
    time = {
      source  = "hashicorp/time"
      version = "= 0.13.1"
    }
  }
}

resource "xcsh_namespace" "test" {
  name = "example"
}

resource "time_sleep" "wait_for_namespace" {
  depends_on      = [xcsh_namespace.test]
  create_duration = "5s"
}

resource "xcsh_malicious_user_mitigation" "test" {
  depends_on  = [time_sleep.wait_for_namespace]
  name        = "example-value"
  namespace   = xcsh_namespace.test.name
  description = "Test malicious user mitigation with all attributes"
  disable     = false

  labels = {
    environment = "test"
    team        = "security"
  }

  annotations = {
    purpose = "testing"
  }
}
```

<a id="canonical-1021333202112011-1013203201111331-1231122032023311-3131300231110212-2231232021112203-2020100100110122-0320332111201312-2313030011321112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Block example

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md#canonical-1101131202021321-3221321112031101-3221121031010113-2211012203220000-0302111121132132-3121301122233133-1013203002033303-3311201201113123)
- [Examples](resources--malicious_user_mitigation--examples--group-001.md#canonical-0122333121121121-2133233012323031-2100322101123321-3101010121133222-1002313333123011-2132000233033313-0320000203211130-1132123110011023)
- Block

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_malicious_user_mitigation/block.tf`; digest `sha256:43058f9bf900f0f24a9521e75d8a2c8d32a8799733304a23a757ce55a8c8ffef`.

```terraform
# Block — Acceptance-test-derived Configuration
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

resource "xcsh_malicious_user_mitigation" "test" {
  name      = "example"
  namespace = "system"

  mitigation_type {
    rules {
      threat_level {
        high = {}
      }
      mitigation_action {
        block_temporarily = {}
      }
    }
  }
}
```

<a id="canonical-0032002003213131-3101211313010231-1223010023132031-2211303333320030-2202213012323012-1301021111132001-1131030212123103-2100312220110012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Captcha example

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md#canonical-1101131202021321-3221321112031101-3221121031010113-2211012203220000-0302111121132132-3121301122233133-1013203002033303-3311201201113123)
- [Examples](resources--malicious_user_mitigation--examples--group-001.md#canonical-0122333121121121-2133233012323031-2100322101123321-3101010121133222-1002313333123011-2132000233033313-0320000203211130-1132123110011023)
- Captcha

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_malicious_user_mitigation/captcha.tf`; digest `sha256:900a9ff0aef2add7bba28c6bef38ed846bf7b633de41900982a3c7d386dce1a1`.

```terraform
# Captcha — Acceptance-test-derived Configuration
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

resource "xcsh_malicious_user_mitigation" "test" {
  name      = "example"
  namespace = "system"

  mitigation_type {
    rules {
      threat_level {
        high = {}
      }
      mitigation_action {
        captcha_challenge = {}
      }
    }
  }
}
```

<a id="canonical-3213332111110303-0233120032031331-3130132222023022-3233112232230001-1110123013010013-2020020120033103-1130312223101333-3203101321203012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Js challenge example

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md#canonical-1101131202021321-3221321112031101-3221121031010113-2211012203220000-0302111121132132-3121301122233133-1013203002033303-3311201201113123)
- [Examples](resources--malicious_user_mitigation--examples--group-001.md#canonical-0122333121121121-2133233012323031-2100322101123321-3101010121133222-1002313333123011-2132000233033313-0320000203211130-1132123110011023)
- Js challenge

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_malicious_user_mitigation/js-challenge.tf`; digest `sha256:31cce82837db63f5b4c06f4195de3f808d4789b7437b6ae8148fa42b9b7c4194`.

```terraform
# JsChallenge — Acceptance-test-derived Configuration
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

resource "xcsh_malicious_user_mitigation" "test" {
  name      = "example"
  namespace = "system"

  mitigation_type {
    rules {
      threat_level {
        medium = {}
      }
      mitigation_action {
        javascript_challenge = {}
      }
    }
  }
}
```

<a id="canonical-3301001202201031-1213330213031310-1333230003311100-2013300321320202-0113311010311321-0233322120202133-3020120221033023-3023220011010101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md#canonical-1101131202021321-3221321112031101-3221121031010113-2211012203220000-0302111121132132-3121301122233133-1013203002033303-3311201201113123)
- [Examples](resources--malicious_user_mitigation--examples--group-001.md#canonical-0122333121121121-2133233012323031-2100322101123321-3101010121133222-1002313333123011-2132000233033313-0320000203211130-1132123110011023)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_malicious_user_mitigation/resource.tf`; digest `sha256:cb208f65f7f8d9724e94da5a2c35bd650a88bc8059b0aeced4b7db40e67c44ff`.

```terraform
# MaliciousUserMitigation Resource Example
# Manages malicious_user_mitigation creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic MaliciousUserMitigation configuration
resource "xcsh_malicious_user_mitigation" "example" {
  name      = "example-malicious-user-mitigation"
  namespace = "staging"
}
```

<a id="canonical-0332011323203221-2212021231330001-3022103131310300-1030022012220312-0103300311201021-3223103001302132-0300002300321220-3012322030320300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## With annotations example

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md#canonical-1101131202021321-3221321112031101-3221121031010113-2211012203220000-0302111121132132-3121301122233133-1013203002033303-3311201201113123)
- [Examples](resources--malicious_user_mitigation--examples--group-001.md#canonical-0122333121121121-2133233012323031-2100322101123321-3101010121133222-1002313333123011-2132000233033313-0320000203211130-1132123110011023)
- With annotations

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_malicious_user_mitigation/with-annotations.tf`; digest `sha256:10af11e1fcedf0c1a38131c01416f578defc093012b6a6669770dd539168e15b`.

```terraform
# WithAnnotations — Acceptance-test-derived Configuration
# Extracted from an acceptance test helper.
# No new live API validation is claimed.

terraform {
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
    time = {
      source  = "hashicorp/time"
      version = "= 0.13.1"
    }
  }
}

resource "xcsh_namespace" "test" {
  name = "example"
}

resource "time_sleep" "wait_for_namespace" {
  depends_on      = [xcsh_namespace.test]
  create_duration = "5s"
}

resource "xcsh_malicious_user_mitigation" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = "example-value"
  namespace  = xcsh_namespace.test.name

  annotations = {
    example-key = "example-value"
  }
}
```

<a id="canonical-0122001301220132-0333321133231321-2301323122121103-2200220230231102-1321331121211023-1120012130313133-2321011220121032-1300222031323120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## With description example

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md#canonical-1101131202021321-3221321112031101-3221121031010113-2211012203220000-0302111121132132-3121301122233133-1013203002033303-3311201201113123)
- [Examples](resources--malicious_user_mitigation--examples--group-001.md#canonical-0122333121121121-2133233012323031-2100322101123321-3101010121133222-1002313333123011-2132000233033313-0320000203211130-1132123110011023)
- With description

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_malicious_user_mitigation/with-description.tf`; digest `sha256:a0561f064cef4a0bf4898a53eed03d3f77741d1f6b9b0474fd171ea08dcd9202`.

```terraform
# WithDescription — Acceptance-test-derived Configuration
# Extracted from an acceptance test helper.
# No new live API validation is claimed.

terraform {
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
    time = {
      source  = "hashicorp/time"
      version = "= 0.13.1"
    }
  }
}

resource "xcsh_namespace" "test" {
  name = "example"
}

resource "time_sleep" "wait_for_namespace" {
  depends_on      = [xcsh_namespace.test]
  create_duration = "5s"
}

resource "xcsh_malicious_user_mitigation" "test" {
  depends_on  = [time_sleep.wait_for_namespace]
  name        = "example-value"
  namespace   = xcsh_namespace.test.name
  description = "example-description"
}
```

<a id="canonical-3020323201221023-0013011020311132-0312111323302223-3333233023023120-0303012131030212-0213211312221323-1101020303111311-2212033020133012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## With labels example

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md#canonical-1101131202021321-3221321112031101-3221121031010113-2211012203220000-0302111121132132-3121301122233133-1013203002033303-3311201201113123)
- [Examples](resources--malicious_user_mitigation--examples--group-001.md#canonical-0122333121121121-2133233012323031-2100322101123321-3101010121133222-1002313333123011-2132000233033313-0320000203211130-1132123110011023)
- With labels

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_malicious_user_mitigation/with-labels.tf`; digest `sha256:05d2d99ea2c9b4398348e89d32380bd3c6c3cf7e2154f51c8b74dcc85481aec2`.

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
    time = {
      source  = "hashicorp/time"
      version = "= 0.13.1"
    }
  }
}

resource "xcsh_namespace" "test" {
  name = "example"
}

resource "time_sleep" "wait_for_namespace" {
  depends_on      = [xcsh_namespace.test]
  create_duration = "5s"
}

resource "xcsh_malicious_user_mitigation" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = "example-value"
  namespace  = xcsh_namespace.test.name

  labels = {
    example-key = "example-value"
  }
}
```

<a id="canonical-3202001211201031-1231330122022303-3311223013133003-2000101313231321-2130030101220033-0132303322101130-2333223130131021-1111222001302220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## With mitigation type example

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md#canonical-1101131202021321-3221321112031101-3221121031010113-2211012203220000-0302111121132132-3121301122233133-1013203002033303-3311201201113123)
- [Examples](resources--malicious_user_mitigation--examples--group-001.md#canonical-0122333121121121-2133233012323031-2100322101123321-3101010121133222-1002313333123011-2132000233033313-0320000203211130-1132123110011023)
- With mitigation type

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_malicious_user_mitigation/with-mitigation-type.tf`; digest `sha256:14ca67cc624bf1e714701869d9bfc825ab0fb5591e24e35f489dcb0a69680563`.

```terraform
# WithMitigationType — Acceptance-test-derived Configuration
# Extracted from an acceptance test helper.
# No new live API validation is claimed.

terraform {
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
    time = {
      source  = "hashicorp/time"
      version = "= 0.13.1"
    }
  }
}

resource "xcsh_namespace" "test" {
  name = "example"
}

resource "time_sleep" "wait_for_namespace" {
  depends_on      = [xcsh_namespace.test]
  create_duration = "5s"
}

resource "xcsh_malicious_user_mitigation" "test" {
  depends_on  = [time_sleep.wait_for_namespace]
  name        = "example-value"
  namespace   = xcsh_namespace.test.name
  description = "Malicious user mitigation with mitigation type configuration"

  mitigation_type {
    rules {
      threat_level {
        high = {}
      }
      mitigation_action {
        block_temporarily = {}
      }
    }
  }
}
```
