---
page_title: "xcsh_virtual_host reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_virtual_host reference."
---

# xcsh_virtual_host reference

<a id="canonical-2013230120032311-0000310023210312-2000120021231311-1320012302201130-1100333011330310-2301201021033112-2111013212322221-0210110110221022"></a>

## `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-1321321311010002-0333130122331222-0233010210103222-1020300033133113-0220033020222102-2031100330302111-3223213023203230-1020203123301203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.common_params.tls_certificates.use_system_defaults` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [tls_parameters](resources--virtual_host--reference--group-002.md#canonical-0331211001213000-2022200002023320-0333201001301110-2110111020031121-0112102012231123-2103133133121332-1220021100301010-3112201011023323)
- [tls_parameters.common_params](resources--virtual_host--reference--group-002.md#canonical-0333211303112320-2232312331020330-3200001211332113-1332003022321233-0123311023033203-1001103201023313-1200233332212321-3130010111203221)
- [tls_parameters.common_params.tls_certificates](resources--virtual_host--reference--group-002.md#canonical-0200003231223022-1321323211123233-0123321102300103-1301210333113212-1111032300001213-3313322300133213-3311332031202333-2110202132233203)
- tls_parameters.common_params.tls_certificates.use_system_defaults

<a id="canonical-3233303122201020-3302201102012022-2002230132202021-1022313312101123-2320113202122202-1223303101311133-2323210201310212-0132132113032203"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for use system defaults.

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
use_system_defaults = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1100132303312233-2302331223202323-2232231202112111-2233122222201102-3320121233323003-0022333231032202-1210011013012132-1022023001212123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.common_params.validation_params` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [tls_parameters](resources--virtual_host--reference--group-002.md#canonical-0331211001213000-2022200002023320-0333201001301110-2110111020031121-0112102012231123-2103133133121332-1220021100301010-3112201011023323)
- [tls_parameters.common_params](resources--virtual_host--reference--group-002.md#canonical-0333211303112320-2232312331020330-3200001211332113-1332003022321233-0123311023033203-1001103201023313-1200233332212321-3130010111203221)
- tls_parameters.common_params.validation_params

<a id="canonical-2031103232112023-1301131322123321-3322010330112122-2333130021322233-1130323231323001-3333003021210333-3103131313200010-2310321101122310"></a>

Type: `"object"`. single nested block, Optional.

This includes URL for a trust store, whether SAN verification is required and list of Subject Alt
Names for verification.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]"
}
```

Terraform syntax:

```terraform
validation_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-2301330301201003-2010333331220223-1122331200131302-0330033230121221-1002231320303311-0232101211331100-2212320022001213-1320013320012103"></a>

### Direct properties for `tls_parameters.common_params.validation_params`

<a id="canonical-2300021102323300-0221100330002100-3333213200321211-2031123132031202-3113301010112003-0223102333201011-3303311102310130-2213332022212001"></a>

#### `tls_parameters.common_params.validation_params.skip_hostname_verification` property

Type: `"bool"`. Optional.

When True, skip verification of hostname i.e. CN/Subject Alt Name of certificate is not matched to
the connecting hostname.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [trusted_ca](resources--virtual_host--reference--group-003.md#canonical-2121201203020333-2300231002003323-3130133031231300-2232003010111033-3233213301303021-1030121101132021-2100030202322021-2020223201311311): complete subsection reference.

<a id="canonical-2201330001003130-1103011331220233-0222132032022310-0202130110120123-3313013301131232-2100012011220310-2311130330033130-0310300011033013"></a>

<a id="canonical-1212201002222100-2022321301323131-2123323321301213-1233111111123230-0022100312332211-3330112220122321-0100321101132330-3131320012113120"></a>

#### `tls_parameters.common_params.validation_params.trusted_ca_url` property

Type: `"string"`. Optional.

Exclusive with \[trusted\_ca\] Inline Root CA Certificate.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

<a id="canonical-2322030231212021-2033120211323221-0300330333313200-2030313333202300-0101103133202310-2102020323122011-3120212302200022-0213121113110013"></a>

<a id="canonical-1112033323203213-1230302022330223-0110030132110023-1111131030022202-3102302332323321-2332003320020303-1311212000203230-3320020023130300"></a>

#### `tls_parameters.common_params.validation_params.verify_subject_alt_names` property

Type: `["list", "string"]`. Optional.

List of acceptable Subject Alt Names/CN in the peer's certificate. When skip\_hostname\_verification
is false and verify\_subject\_alt\_names is empty, the hostname of the peer will be used for
matching against SAN/CN of peer's certificate.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2121201203020333-2300231002003323-3130133031231300-2232003010111033-3233213301303021-1030121101132021-2100030202322021-2020223201311311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.common_params.validation_params.trusted_ca` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [tls_parameters](resources--virtual_host--reference--group-002.md#canonical-0331211001213000-2022200002023320-0333201001301110-2110111020031121-0112102012231123-2103133133121332-1220021100301010-3112201011023323)
- [tls_parameters.common_params](resources--virtual_host--reference--group-002.md#canonical-0333211303112320-2232312331020330-3200001211332113-1332003022321233-0123311023033203-1001103201023313-1200233332212321-3130010111203221)
- [tls_parameters.common_params.validation_params](resources--virtual_host--reference--group-003.md#canonical-1100132303312233-2302331223202323-2232231202112111-2233122222201102-3320121233323003-0022333231032202-1210011013012132-1022023001212123)
- tls_parameters.common_params.validation_params.trusted_ca

<a id="canonical-2102002130033323-2330120222232233-0302302223102230-2321122000230332-2002221220220110-0222020122210222-0310213113000300-1221310213031203"></a>

Type: `"object"`. single nested block, Optional.

Root CA Certificate Reference. Reference to Root CA Certificate.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
trusted_ca {
  # Configure direct properties listed below.
}
```

<a id="canonical-2312013010303301-2311133022100130-1313323120230213-1021302022201120-0002213003322002-1122022203321230-0312323322301001-3212231330212133"></a>

### Direct properties for `tls_parameters.common_params.validation_params.trusted_ca`

- [trusted_ca_list](resources--virtual_host--reference--group-003.md#canonical-0003333010322121-3122312223300132-0221230312011321-3031120011203202-2020200121301130-3030230011200310-3113333031131002-3010201203110311): complete subsection reference.

<a id="canonical-0003333010322121-3122312223300132-0221230312011321-3031120011203202-2020200121301130-3030230011200310-3113333031131002-3010201203110311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [tls_parameters](resources--virtual_host--reference--group-002.md#canonical-0331211001213000-2022200002023320-0333201001301110-2110111020031121-0112102012231123-2103133133121332-1220021100301010-3112201011023323)
- [tls_parameters.common_params](resources--virtual_host--reference--group-002.md#canonical-0333211303112320-2232312331020330-3200001211332113-1332003022321233-0123311023033203-1001103201023313-1200233332212321-3130010111203221)
- [tls_parameters.common_params.validation_params](resources--virtual_host--reference--group-003.md#canonical-1100132303312233-2302331223202323-2232231202112111-2233122222201102-3320121233323003-0022333231032202-1210011013012132-1022023001212123)
- [tls_parameters.common_params.validation_params.trusted_ca](resources--virtual_host--reference--group-003.md#canonical-2121201203020333-2300231002003323-3130133031231300-2232003010111033-3233213301303021-1030121101132021-2100030202322021-2020223201311311)
- tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list

<a id="canonical-3012330211123302-1022033223131232-3113120133013000-0000221203203123-3210202130332330-3210013003001012-2202232013131322-1232222121013122"></a>

Type: `"object"`. list nested block, Optional.

Root CA Certificate Reference. Reference to Root CA Certificate.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

Terraform syntax:

```terraform
trusted_ca_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1012311322021200-2303100021210313-2103201323121001-0132100022001021-3320023032230102-0030022302113321-1000032333011131-2201330120033110"></a>

### Direct properties for `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list`

<a id="canonical-3212002232220301-3313010002121021-2003113122131223-2111300310233020-0022220303032011-2230022111021101-0120003113132301-2130223012221233"></a>

#### `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2313110201103221-0131000112102213-1323231013021300-1022212013303212-0201001231110210-1231003332302301-1131020200010130-2012331321330100"></a>

<a id="canonical-0321120100131033-0312333221001230-1002232130003132-3213212021022001-2003312301003321-1233120021031121-2320203123103233-1313032022111312"></a>

#### `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1023023233122030-3330233203231330-0132100103311113-3031010113003331-0312120100203022-1201031103120330-0111221121300102-2320213000002233"></a>

<a id="canonical-3201122101002120-1223101323123220-2212023131203301-2213133301101033-1323132311210312-0023223123110000-1313010120300000-1330021331321331"></a>

#### `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3031210003122301-1210322301302202-0212313003122220-2113312132330000-2020123000222303-1121011311110210-2103232030313003-3001122103023322"></a>

<a id="canonical-1312331202300222-2000020130310300-1132123030003222-2111003032133202-1321213221001102-2033222312123210-0001332033222002-2231330033211232"></a>

#### `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3323031231011300-2221021132130121-0033000333033011-3100330102112110-1332121030013310-3113232123233002-2023332321103121-1120321321121200"></a>

<a id="canonical-1312333010220131-3212320132013301-3013220330321131-0100031123010022-3100331301333323-1300330230002313-0121210003001223-3120102211022113"></a>

#### `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.uid` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2213212201113300-0223221321112131-2300013131310300-2330130313033001-3003001223303232-3000310231222000-2101120210302122-2300231020001031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.no_client_certificate` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [tls_parameters](resources--virtual_host--reference--group-002.md#canonical-0331211001213000-2022200002023320-0333201001301110-2110111020031121-0112102012231123-2103133133121332-1220021100301010-3112201011023323)
- tls_parameters.no_client_certificate

<a id="canonical-2323112302311330-1200313313203222-0323031221031333-3111020121011033-0233231302132131-3303333311000121-0000010300221333-2213020222010031"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
no_client_certificate = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1012201010031002-2232101323213000-1022110300231232-0230000221021101-1203021223001011-2301313113020120-2002021332022003-0311230002232310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `user_identification` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- user_identification

<a id="canonical-3123231102032223-3202112001213220-2323200123223203-3121212301113113-1210231333210211-1123211210111022-3223021311311112-0221103332030102"></a>

Type: `"object"`. list nested block, Optional.

A reference to user\_identification object. The rules in the user\_identification object are
evaluated to determine the user identifier to be rate limited.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

Terraform syntax:

```terraform
user_identification {
  # Configure direct properties listed below.
}
```

<a id="canonical-0332211313222013-0200232320233113-1131232210213221-0000320120132131-1320210130031332-3120001201330233-0332220210200001-0102303032203231"></a>

### Direct properties for `user_identification`

<a id="canonical-0120222111310222-0320333030121310-1102300020313231-1030130301033323-3221101111030120-1212213333213300-1203123010003020-0211031003300033"></a>

#### `user_identification.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2120122131223012-0223301301103212-3223131020130310-0232232310300032-2221002122032323-3232211101011112-0231223121212011-3211100213131103"></a>

<a id="canonical-0211001322130020-0223112222323112-0322200123203322-1330222020212232-2003210101001103-3211010300123202-3202310121322220-3211100131333301"></a>

#### `user_identification.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0203221000130021-1111003301123201-1103331311111102-2111332110120302-3102113223330131-1213123230022303-1312333332002220-3203320022133102"></a>

<a id="canonical-1121230102113101-0100010110230200-0101003330311023-0123213003302000-3100310300131001-0332311001010201-2003133321321010-1031133323223032"></a>

#### `user_identification.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1101222132001233-2111020200313002-1001210010013223-0010312130211201-0113130311031323-0311031031232233-2133011202013233-2210300023310220"></a>

<a id="canonical-2032123311021011-2210222233223110-2030330120031102-3323102100330200-0232302032023113-2320021331002123-0001133302313211-0123103000333010"></a>

#### `user_identification.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0322121032330201-3201023101112312-2101102133132312-3332212232010300-2201212200202233-1002222210023032-2021100121233103-0332130232003102"></a>

<a id="canonical-2232233123333021-1331031302121101-1312120123131302-3023000201010203-3201120033023312-0031210311212130-0102233032233230-0132223220032332"></a>

#### `user_identification.uid` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2322100101202101-2331131111122123-0200313013302211-3200032010130132-1310103000220023-2103022322122223-2311300301002030-1022212031023331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_type` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- waf_type

<a id="canonical-3230013003023000-1033211112313122-3023101303102010-3313122122303000-1202210222021131-1031330201322010-0303332312111121-3302321220210320"></a>

Type: `"object"`. single nested block, Optional.

WAF instance will be pointing to an app\_firewall object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ref_type": "[\"app_firewall\",\"disable_waf\",\"inherit_waf\"]"
}
```

Terraform syntax:

```terraform
waf_type {
  # Configure direct properties listed below.
}
```

<a id="canonical-2110211020000103-1220010130012003-0021003033221213-0001002220132320-1001303212300320-2201303103222030-1332102333211010-2010020330120322"></a>

### Direct properties for `waf_type`

- [app_firewall](resources--virtual_host--reference--group-003.md#canonical-3013121233031313-3131021232212222-3222231213232301-1333021332220013-1002103010033211-0110331022300302-1330321011210331-2131112331011031): complete subsection reference.

- [disable_waf](resources--virtual_host--reference--group-003.md#canonical-0200013230100212-1300201000323321-0100110000112200-1102021123321212-1003031221312331-3213322010301212-3122122332032230-2231003012221103): complete subsection reference.

- [inherit_waf](resources--virtual_host--reference--group-003.md#canonical-1310320202133212-2221033303020302-0011200030000030-2010112132322220-1232023133223203-1113301212012300-1132200001133223-1303021333223200): complete subsection reference.

<a id="canonical-3013121233031313-3131021232212222-3222231213232301-1333021332220013-1002103010033211-0110331022300302-1330321011210331-2131112331011031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_type.app_firewall` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [waf_type](resources--virtual_host--reference--group-003.md#canonical-2322100101202101-2331131111122123-0200313013302211-3200032010130132-1310103000220023-2103022322122223-2311300301002030-1022212031023331)
- waf_type.app_firewall

<a id="canonical-3213022012121222-0002330131002030-3310033233203333-1033330211110330-2120011000001330-2201021020201233-0332123000320320-2212020330301333"></a>

Type: `"object"`. single nested block, Optional.

A list of references to the app\_firewall configuration objects.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
app_firewall {
  # Configure direct properties listed below.
}
```

<a id="canonical-1123120203232203-2210223222021301-1012023303112211-0210320321301031-0321022130322213-0333201023100002-0011033330020111-2023211002211332"></a>

### Direct properties for `waf_type.app_firewall`

- [app_firewall](resources--virtual_host--reference--group-003.md#canonical-3002131032323211-1313323003223001-1030222330120001-3220101010122031-0213113222211303-3312312200323101-1132121102031200-3312230200131203): complete subsection reference.

<a id="canonical-3002131032323211-1313323003223001-1030222330120001-3220101010122031-0213113222211303-3312312200323101-1132121102031200-3312230200131203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_type.app_firewall.app_firewall` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [waf_type](resources--virtual_host--reference--group-003.md#canonical-2322100101202101-2331131111122123-0200313013302211-3200032010130132-1310103000220023-2103022322122223-2311300301002030-1022212031023331)
- [waf_type.app_firewall](resources--virtual_host--reference--group-003.md#canonical-3013121233031313-3131021232212222-3222231213232301-1333021332220013-1002103010033211-0110331022300302-1330321011210331-2131112331011031)
- waf_type.app_firewall.app_firewall

<a id="canonical-2123223303123331-0321031133303331-2022221030110212-0301330033003311-1012333203333123-3230110301020030-1221011102120001-0210201321321333"></a>

Type: `"object"`. list nested block, Optional.

References to an Application Firewall configuration object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.num_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.num_items": "1"
  }
}
```

Terraform syntax:

```terraform
app_firewall {
  # Configure direct properties listed below.
}
```

<a id="canonical-0211133231112321-0023232023013213-0003300030311201-2133301310033213-0300111113123023-3223212030320322-2233211000220001-0200300223222222"></a>

### Direct properties for `waf_type.app_firewall.app_firewall`

<a id="canonical-0010032012013331-0303102231333212-3201133131010023-3133001021110200-2202200133000320-0130230102311131-1323223223130000-2101322301322200"></a>

#### `waf_type.app_firewall.app_firewall.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1200031333031012-0023302213113102-1121212310312203-3310320333231022-0321110113222010-2030302020120332-1303232230333001-1010322323123100"></a>

<a id="canonical-2202330021332300-2021330221011131-1312012033001001-1313113022331131-2113203223003011-3221032011221231-1231131230022213-0110000021032112"></a>

#### `waf_type.app_firewall.app_firewall.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2031123111321331-3332320323132012-3232332013331211-1212331111033220-0003212320322001-1121312120331231-1222212232220321-2313321203022032"></a>

<a id="canonical-1331033322311300-0102023130121032-0223211231033011-1221200121310123-3010300220323230-3311313013232112-1220320213323303-2222112221110232"></a>

#### `waf_type.app_firewall.app_firewall.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3101322222023123-0001313303001031-0033203333000220-0003233001000020-2212202003310100-1312320202023303-1201312112203202-0131222003332121"></a>

<a id="canonical-0123232222000021-1101213101210211-3020211313231100-1202131101203221-1111223031200311-3103011131223222-2320331003331033-1113010233023200"></a>

#### `waf_type.app_firewall.app_firewall.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1123303212111101-0231123003302020-1031321303000110-3122101313300322-0030111332130222-1221123331013331-2012120220012013-3332131100221002"></a>

<a id="canonical-1322030031123122-3101301231102223-3123113011120003-3110030103013133-3232301220002010-1003033301020311-3010230333222332-3312322033122331"></a>

#### `waf_type.app_firewall.app_firewall.uid` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0200013230100212-1300201000323321-0100110000112200-1102021123321212-1003031221312331-3213322010301212-3122122332032230-2231003012221103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_type.disable_waf` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [waf_type](resources--virtual_host--reference--group-003.md#canonical-2322100101202101-2331131111122123-0200313013302211-3200032010130132-1310103000220023-2103022322122223-2311300301002030-1022212031023331)
- waf_type.disable_waf

<a id="canonical-0302310212030013-0311333130001103-0122210310011201-1201023311102112-0230033112133120-3001212100112231-3123023021221320-0220210212310210"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable waf.

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
disable_waf = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1310320202133212-2221033303020302-0011200030000030-2010112132322220-1232023133223203-1113301212012300-1132200001133223-1303021333223200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_type.inherit_waf` properties

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [waf_type](resources--virtual_host--reference--group-003.md#canonical-2322100101202101-2331131111122123-0200313013302211-3200032010130132-1310103000220023-2103022322122223-2311300301002030-1022212031023331)
- waf_type.inherit_waf

<a id="canonical-0011131120230110-2322330020322312-1011301301130303-1223322210131003-0221300000132021-1302213300322220-0231320323010231-2321222020130331"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for inherit waf.

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
inherit_waf = {}
```

This is an empty object or choice marker. It has no direct properties.
