---
page_title: "xcsh_udp_loadbalancer reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_udp_loadbalancer reference."
---

# xcsh_udp_loadbalancer reference

<a id="canonical-2100233330113013-1011020101321013-1113020011303120-3332020200220223-1130023001010003-0211031020201213-0121011023231220-3020101330032130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3101323131223122-1113200230210332-1231011221101232-0012001020312122-2101031112000000-2223223013330133-1131222121210021-0023131203102200"></a>

## Property reference — Property reference / 012202002322 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)
- Property reference

<a id="canonical-1210220021030012-3233330330000123-0330321000102023-3110200232001231-3101132200010133-1332001313233102-1210130000303031-1020321010313001"></a>

## Direct properties — Property reference / 012202002322 / 3

- [active_service_policies](data-sources--udp_loadbalancer--reference--group-001.md#canonical-3012132102203030-3133031323022332-1202102030001302-2101121232121222-1230212330110021-3133230213123232-3113221120113321-3020002310301231): complete subsection reference.

- [advertise_custom](data-sources--udp_loadbalancer--reference--group-001.md#canonical-0102011322232323-2032232030221030-3013012213003033-1131020131301123-2120010001223303-0101100312231222-3210322122303222-2321103220212112): complete subsection reference.

- [advertise_on_public](data-sources--udp_loadbalancer--reference--group-001.md#canonical-3111011220331032-2313033231112301-0320230202101033-3101213131331212-0221323210201200-3230133203131100-3011303013112000-3303002032321131): complete subsection reference.

- [advertise_on_public_default_vip](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2311322310303223-0231020022011110-0310223232001233-0311013120223030-0200302330003100-2330123303103011-3100022130033000-2112131213103102): complete subsection reference.

<a id="canonical-3310231033303130-2223321111200202-0023002323121012-1001212300302022-3030301202212131-0231230123212333-3303013211031222-1333202200302321"></a>

<a id="canonical-0333122333300132-2202331023131333-3010132033333032-2001210102222201-0102333001111011-1231323130113000-1010333113311320-1113103313203032"></a>

## annotations property — Property reference / 012202002322 / 4

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Upstream description:

Annotations is an unstructured key-value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

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
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-3000313021223332-1220013131010233-2230023010231310-1333333330320333-3302232002332101-2203200120311022-3132102010233010-3103111233310303"></a>

<a id="canonical-0302111201333212-2233120031102133-3012222313132111-3201122220023230-0101101001131203-1111100221132020-2313210003021321-1321111221311320"></a>

## description property — Property reference / 012202002322 / 5

Type: `"string"`. Computed.

Description of the UDPLoadBalancer.

Upstream description:

Human readable description for the object.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1200,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1200
    },
    "category": "discovery",
    "characterSet": {
      "description": "Free text with UTF-8 support"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1200,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  }
}
```

<a id="canonical-3323033033220322-2031333033002211-0331002300030032-0302210300302001-2010323221231102-2222330333003121-3232100103223012-1203203303313200"></a>

<a id="canonical-3023321332111133-1300102331122322-1202212011032210-3302112023331303-2322110323311220-2020021012022023-1000121113321003-3123210323110132"></a>

## dns_volterra_managed property — Property reference / 012202002322 / 6

Type: `"bool"`. Computed.

DNS records for domains will be managed automatically by F5 Distributed Cloud. As a prerequisite,
the domain to be delegated to F5 Distributed Cloud using the Delegated Domain feature or a DNS CNAME
record must be created in your DNS provider's portal.

Upstream description:

DNS records for domains will be managed automatically by F5 Distributed Cloud. As a prerequisite,
the domain to be delegated to F5 Distributed Cloud using the Delegated Domain feature or a DNS CNAME
record must be created in your DNS provider's portal.

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

- [do_not_advertise](data-sources--udp_loadbalancer--reference--group-001.md#canonical-0131000331300322-0103301113212111-2211110212313310-3002002230210313-3103111021003321-1003003123210201-3122100201032330-0222313230310031): complete subsection reference.

<a id="canonical-1220132012302321-0320120103002132-2121222003133022-1202210011223300-3200221120002220-2222221230220010-3021101213000131-1311110130020210"></a>

<a id="canonical-1003313331213322-0203121012203023-0132033123302233-0112133302230200-2303010330121012-2121313021330020-1200313231323122-0220331111220220"></a>

## domains property — Property reference / 012202002322 / 7

Type: `["list", "string"]`. Computed.

List of domains (host/authority header) that will be matched to this load balancer.

Upstream description:

A list of domains (host/authority header) that will be matched to this load balancer.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.hostname": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.hostname": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [hash_policy_choice_random](data-sources--udp_loadbalancer--reference--group-001.md#canonical-3002201000313203-2123320202312030-0022330303110220-0230302222121120-3201231231321301-0232331300203223-0022233321021130-0002331230323021): complete subsection reference.

- [hash_policy_choice_round_robin](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1223232202021103-2020233233211000-0201322313330131-2301121001023022-2021122103113111-2320302310333212-3013001031012230-2333021133101031): complete subsection reference.

- [hash_policy_choice_source_ip_stickiness](data-sources--udp_loadbalancer--reference--group-001.md#canonical-0312203212031200-0300032003330222-3023010030101230-1213021131301013-0011220230203121-3021031233331130-0320012111103302-1133302313101320): complete subsection reference.

<a id="canonical-3213311023103102-1313122013313231-1222220121211033-3333001323230220-1010032110332001-0032031102032223-0132232200210310-1230231023112313"></a>

<a id="canonical-0112013010121033-3202330212013200-2312220200002311-2333332113300103-0022131210021321-1231033233023210-1032013120310003-3323322130110213"></a>

## ID property — Property reference / 012202002322 / 8

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0011213002103221-2110330321003022-1031132023330022-3113000222313032-0220200222012330-1133013233331223-3311021203122330-0031132330010100"></a>

<a id="canonical-0133311102213321-2101303310022020-3033312013100122-0231321322033331-1033011311330102-2121123020010130-2113300213302300-1003113130311310"></a>

## idle_timeout property — Property reference / 012202002322 / 9

Type: `"number"`. Computed.

The amount of time that a session can exist without upstream or downstream activity, in
milliseconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 30000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "30000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "30000"
  }
}
```

<a id="canonical-1022302202222302-0312132100003212-2333131133200332-0230102122103120-0231323013301303-1311320220033112-1222133310103311-3211312130102232"></a>

<a id="canonical-3021030033202302-2220310210200320-1002032333200222-1203020212333213-3323103223111021-0032232220000220-0320313100112323-3313033220132203"></a>

## labels property — Property reference / 012202002322 / 10

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

Upstream description:

Map of string keys and values that can be used to organize and categorize (scope and select) objects
as chosen by the user. Values specified here will be used by selector expression.

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

<a id="canonical-0201211311302131-0033213301132110-3130311012101302-1113311023232122-3230023121311030-1332222300233120-2333303030131113-3223211023313010"></a>

<a id="canonical-3313121233333021-2033000223001333-1020202031123323-3112333130032310-2111322300331012-3121330033233013-2221302032110310-3123101112023213"></a>

## listen_port property — Property reference / 012202002322 / 11

Type: `"number"`. Computed.

\[OneOf: listen\_port, port\_ranges\] Exclusive with \[port\_ranges\] Listen Port for this load
balancer.

Upstream description:

Exclusive with \[port\_ranges\] Listen Port for this load balancer.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

OneOf alternatives in this subsection:

- [listen_port](data-sources--udp_loadbalancer--reference--group-001.md#canonical-0201211311302131-0033213301132110-3130311012101302-1113311023232122-3230023121311030-1332222300233120-2333303030131113-3223211023313010)
- [port_ranges](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2032300322031010-2200011221101230-3213232331000322-1032022312031312-2022122301233121-3000300302311221-3322033023022132-3103220223211103)

Select alternatives according to the provider validators above.

<a id="canonical-0000021102112212-2221111231031123-1302031222220131-0011223121301113-0231310113103100-2001103212303121-0232131222233002-1113212102321313"></a>

<a id="canonical-2302113330203031-1322220223100233-3222331303333001-0032322302213000-2032303211021233-2231120303230333-3311213133112302-1100212012120233"></a>

## name property — Property reference / 012202002322 / 12

Type: `"string"`. Required.

Name of the UDPLoadBalancer.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
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

<a id="canonical-2223113313130230-3120203122300133-2110110322032132-3221102032030230-1102113012202220-2012113232122330-0223301321110130-3223212302012332"></a>

<a id="canonical-3133320201131202-0221102033121103-0132033003100321-3202003230002323-0200322103122232-2231213032302300-1320102133130011-2221113322323122"></a>

## namespace property — Property reference / 012202002322 / 13

Type: `"string"`. Required.

Namespace where the UDPLoadBalancer exists.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [no_service_policies](data-sources--udp_loadbalancer--reference--group-001.md#canonical-0311102003233330-0121133300020322-3132013220222322-3203312022330021-3033322310112312-2323110002200030-0120301333121123-3011112033313013): complete subsection reference.

- [origin_pools_weights](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2320311033331002-1133320121212213-0320333113110131-2120100213202221-2233020011030301-1322013030212121-1131120221033210-2032330100122213): complete subsection reference.

<a id="canonical-2032300322031010-2200011221101230-3213232331000322-1032022312031312-2022122301233121-3000300302311221-3322033023022132-3103220223211103"></a>

<a id="canonical-0121302011102111-0112100303022003-2211120201322023-1301323300222232-2323202133330213-0202232211123320-1212020021202103-2030112100021333"></a>

## port_ranges property — Property reference / 012202002322 / 14

Type: `"string"`. Computed.

Exclusive with \[listen\_port\] A string containing a comma separated list of port ranges. Each port
range consists of a single port or two ports separated by '-'.

Upstream description:

Exclusive with \[listen\_port\] A string containing a comma separated list of port ranges. Each port
range consists of a single port or two ports separated by "-".

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
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
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  }
}
```

- [service_policies_from_namespace](data-sources--udp_loadbalancer--reference--group-001.md#canonical-3122231333000123-0312111333033323-2323101223110320-3223113220211321-0222023230110012-1103201221010012-0210212333003320-3122113112330300): complete subsection reference.

- [udp](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2320000221111121-2211321333322221-0131123101031120-1011222001302102-3020022130323223-1330300011100301-1023000101320102-3123222213330010): complete subsection reference.

<a id="canonical-3123201113210311-3100232211330100-0030102332032010-3131032102202012-1120200112303303-1303230331032330-2110321221120333-1220100013100211"></a>

## All schema paths — Property reference / 012202002322 / 15

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `active_service_policies` | [active_service_policies](data-sources--udp_loadbalancer--reference--group-001.md#canonical-3022332123102210-1013032320232232-1231133223023010-0131231223312023-0230212023013212-1001032313131122-2233022323231012-2120301003323332) |
| `active_service_policies.policies` | [active_service_policies.policies](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2111200323110000-0011021030120010-1031201133001021-1001332130331211-2233120303330032-1312212330030001-2332131213113232-3202002312301120) |
| `active_service_policies.policies.name` | [active_service_policies.policies.name](data-sources--udp_loadbalancer--reference--group-001.md#canonical-3320111100011310-2001033312312100-1212302033220332-3110222200201312-0232202313211113-1231133133233320-1121310100310100-3303200213201130) |
| `active_service_policies.policies.namespace` | [active_service_policies.policies.namespace](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2032111202021011-0012310030223232-3211020122232320-0012221321231013-0213222210000123-0030102312002122-0313032132313130-3132002312030031) |
| `active_service_policies.policies.tenant` | [active_service_policies.policies.tenant](data-sources--udp_loadbalancer--reference--group-001.md#canonical-3203011200020133-2130212201032113-2023310333111311-1200222220212133-1311301211123020-1013102303102211-2102330211212013-2201010200112100) |
| `advertise_custom` | [advertise_custom](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1100123320322233-0112312301223220-2303322231013321-3113113310222200-2133011023231220-2301331020100210-3313300312303321-2021233211330012) |
| `advertise_custom.advertise_where` | [advertise_custom.advertise_where](data-sources--udp_loadbalancer--reference--group-001.md#canonical-3322330333233133-1301212300100033-1023010031000021-3021130333311021-2201310300222023-2003210122303202-0211023100120231-0211113033310011) |
| `advertise_custom.advertise_where.advertise_dualstack_on_public` | [advertise_custom.advertise_where.advertise_dualstack_on_public](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1030321213303022-1003212332131122-1303000033113112-2023313123223101-3311302103001210-3112102003131213-1221202133220100-3331002233000103) |
| `advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip` | [advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1130020210100000-3220101201001003-0122223222111120-0232122103001221-2130212311233020-1001011112012322-3311120233321203-1230031032203212) |
| `advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.name` | [advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.name](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2321202101022130-1201032203211122-2001222101232303-1130213332232000-1100013333003100-3013110232231002-2111320012212131-1121120033213312) |
| `advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.namespace` | [advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.namespace](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1020023233112120-1221312003112000-3003230211120323-1112330332202333-0022113211233223-0032021323120312-3113200030230001-3122233121021323) |
| `advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.tenant` | [advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.tenant](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2213223330003300-2311332312310010-0330031023012313-1233000121231120-3331222012313133-2212112212002110-1022123131013121-2220013232133100) |
| `advertise_custom.advertise_where.advertise_on_public` | [advertise_custom.advertise_where.advertise_on_public](data-sources--udp_loadbalancer--reference--group-001.md#canonical-0033320301023120-3000102012021310-3111331133303311-0022231210213200-3201110032310220-2132032110330313-3210013312332232-0112032122020013) |
| `advertise_custom.advertise_where.advertise_on_public.public_ip` | [advertise_custom.advertise_where.advertise_on_public.public_ip](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2110111123000211-1331330211201330-2202303103310300-0032001002333122-1031322000330101-1011221001231033-3103212332022131-3033120323233320) |
| `advertise_custom.advertise_where.advertise_on_public.public_ip.name` | [advertise_custom.advertise_where.advertise_on_public.public_ip.name](data-sources--udp_loadbalancer--reference--group-001.md#canonical-0323022032003232-2123332123201120-2112233122213112-1232013333102333-0100133130033203-1311122032130012-1120330011130022-3211330010333010) |
| `advertise_custom.advertise_where.advertise_on_public.public_ip.namespace` | [advertise_custom.advertise_where.advertise_on_public.public_ip.namespace](data-sources--udp_loadbalancer--reference--group-001.md#canonical-3131023030100001-0130221113200100-0302212220013230-2202023022301012-3231321212213013-2220320211323030-0333012230010123-1230220232131112) |
| `advertise_custom.advertise_where.advertise_on_public.public_ip.tenant` | [advertise_custom.advertise_where.advertise_on_public.public_ip.tenant](data-sources--udp_loadbalancer--reference--group-001.md#canonical-0321300331003202-0320332303121011-0210011021002031-2010302301301203-3100031010321003-1012130231323230-0313113321121003-1310230201021112) |
| `advertise_custom.advertise_where.advertise_v6_on_public` | [advertise_custom.advertise_where.advertise_v6_on_public](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2220113100301322-2223200100012202-3110103322221221-1010221000133002-0333132110123033-3121023021102303-2302033010223311-2003302310013332) |
| `advertise_custom.advertise_where.advertise_v6_on_public.public_ip` | [advertise_custom.advertise_where.advertise_v6_on_public.public_ip](data-sources--udp_loadbalancer--reference--group-001.md#canonical-3020231130121203-0003233310222302-0020013031023121-2201100112213010-0102012330310221-0031030020033103-1302030302333303-2323210130303201) |
| `advertise_custom.advertise_where.advertise_v6_on_public.public_ip.name` | [advertise_custom.advertise_where.advertise_v6_on_public.public_ip.name](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1301123013232111-0323022112113331-2020023302320321-1100112000102201-1013001110313010-3300322301100303-2302101232112232-3312123010232131) |
| `advertise_custom.advertise_where.advertise_v6_on_public.public_ip.namespace` | [advertise_custom.advertise_where.advertise_v6_on_public.public_ip.namespace](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1103210133320333-3003010213100021-0300223211021323-1212210302330230-3200213002331213-0300301220032332-0033203300031030-1101120301000011) |
| `advertise_custom.advertise_where.advertise_v6_on_public.public_ip.tenant` | [advertise_custom.advertise_where.advertise_v6_on_public.public_ip.tenant](data-sources--udp_loadbalancer--reference--group-001.md#canonical-3013223323001300-3220120133021032-0210333132020301-1032331332122020-3001312231102310-0311321030303233-0333222131220002-1000101002020200) |
| `advertise_custom.advertise_where.port` | [advertise_custom.advertise_where.port](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2112202310132202-3002213133311331-3223231330103332-3332132002030031-2321213122311100-3202232200113302-1121113212231113-1322130113311010) |
| `advertise_custom.advertise_where.port_ranges` | [advertise_custom.advertise_where.port_ranges](data-sources--udp_loadbalancer--reference--group-001.md#canonical-3323132113031300-3220013131120301-1130030032301301-3113032122232223-3223010312203031-2002100212012311-3213332223023300-3321122210223103) |
| `advertise_custom.advertise_where.site` | [advertise_custom.advertise_where.site](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2212221312130111-0232110003130221-1010331231320201-2200321311120002-3120000021232100-2120130010132133-2222011333033003-2021210321221311) |
| `advertise_custom.advertise_where.site.ip` | [advertise_custom.advertise_where.site.ip](data-sources--udp_loadbalancer--reference--group-001.md#canonical-0111013013022111-0000002233013230-3311100111002020-1103111001021332-2233101123120301-0002112010220303-3321213311210321-3330312022113111) |
| `advertise_custom.advertise_where.site.network` | [advertise_custom.advertise_where.site.network](data-sources--udp_loadbalancer--reference--group-001.md#canonical-0000121300333031-3121213202231132-1100102210022223-0330301302000013-3101322220020231-0002020222303122-2112312032021233-1110101312012100) |
| `advertise_custom.advertise_where.site.site` | [advertise_custom.advertise_where.site.site](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1133221030200030-0000311013101303-3022330032230202-1203032023021322-1333303333023322-2121012233032301-2121231111013310-2101030201231120) |
| `advertise_custom.advertise_where.site.site.name` | [advertise_custom.advertise_where.site.site.name](data-sources--udp_loadbalancer--reference--group-001.md#canonical-0031202101313013-3321223100331100-0132302310101013-0002211122221323-2310023211231013-1220112231232320-2232202023100323-2312120131020333) |
| `advertise_custom.advertise_where.site.site.namespace` | [advertise_custom.advertise_where.site.site.namespace](data-sources--udp_loadbalancer--reference--group-001.md#canonical-3113111133302232-0011011133012331-3000231022021300-1312131331010132-2213202020333313-0121103322111021-1313330032330022-2202020122032220) |
| `advertise_custom.advertise_where.site.site.tenant` | [advertise_custom.advertise_where.site.site.tenant](data-sources--udp_loadbalancer--reference--group-001.md#canonical-3213123021132100-1303223233203333-2321213201231223-1303011300233322-3201031132232222-1201331322222210-2222210302032321-2130131201322011) |
| `advertise_custom.advertise_where.use_default_port` | [advertise_custom.advertise_where.use_default_port](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2321330311032101-3033001320102311-0030323100101321-2012203213002110-2011102032333220-1002332023300233-1320220223333021-0303301113210221) |
| `advertise_custom.advertise_where.virtual_network` | [advertise_custom.advertise_where.virtual_network](data-sources--udp_loadbalancer--reference--group-001.md#canonical-0301200312031021-3021021033202132-1213111301123303-1031113310023222-3103322332311103-0330312223003203-0333111030310133-3031132203131023) |
| `advertise_custom.advertise_where.virtual_network.default_v6_vip` | [advertise_custom.advertise_where.virtual_network.default_v6_vip](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2003113332311120-0003321011000110-3223103020300123-2320200001221012-1222212311130230-3001332323120020-0001311113132300-1213011223333211) |
| `advertise_custom.advertise_where.virtual_network.default_vip` | [advertise_custom.advertise_where.virtual_network.default_vip](data-sources--udp_loadbalancer--reference--group-001.md#canonical-0010220213213022-0013121302323110-0110210311302020-2123323312320210-0230111313011130-3030132020020232-1103311122321112-2221211312221200) |
| `advertise_custom.advertise_where.virtual_network.specific_v6_vip` | [advertise_custom.advertise_where.virtual_network.specific_v6_vip](data-sources--udp_loadbalancer--reference--group-001.md#canonical-0103003003133020-3010331332222230-2330030130013220-0110300013313110-2131113230002103-2023033302222201-3100022300313121-0011210202201000) |
| `advertise_custom.advertise_where.virtual_network.specific_vip` | [advertise_custom.advertise_where.virtual_network.specific_vip](data-sources--udp_loadbalancer--reference--group-001.md#canonical-3212202212220323-3301202332222032-0022311333031220-3301203133133120-0010310201020033-1233331031002122-1000100000312130-1223200302300113) |
| `advertise_custom.advertise_where.virtual_network.virtual_network` | [advertise_custom.advertise_where.virtual_network.virtual_network](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2100120132222112-1333120200020002-3323021130233113-0201032111113232-0023030221111211-0022132322121323-2121011221312031-3021100102101122) |
| `advertise_custom.advertise_where.virtual_network.virtual_network.name` | [advertise_custom.advertise_where.virtual_network.virtual_network.name](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2131000211310123-3023331330303130-0330303302022002-1131123122300100-3312131013232013-3303213001133020-1212033221322303-1310321212103223) |
| `advertise_custom.advertise_where.virtual_network.virtual_network.namespace` | [advertise_custom.advertise_where.virtual_network.virtual_network.namespace](data-sources--udp_loadbalancer--reference--group-001.md#canonical-0122122222321103-3033010301011201-1022011003333021-2121121233103232-2122030112112303-3212032022333301-3330310231303233-0122022100232322) |
| `advertise_custom.advertise_where.virtual_network.virtual_network.tenant` | [advertise_custom.advertise_where.virtual_network.virtual_network.tenant](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1020010113013302-1110120131020231-1301121001111132-3223302032311221-1000332012001102-3002130222130133-3203300003133113-3021120223123303) |
| `advertise_custom.advertise_where.virtual_site` | [advertise_custom.advertise_where.virtual_site](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1122203022113201-1112121013120332-0003203300331102-1232022311123213-3202002211022112-0023020000300013-0121220301323311-2302031110310311) |
| `advertise_custom.advertise_where.virtual_site.network` | [advertise_custom.advertise_where.virtual_site.network](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1213123113003300-2011031302230033-3123030212113130-2010211233323311-1221302303332230-1021122030130203-3132203210312111-2222101132002110) |
| `advertise_custom.advertise_where.virtual_site.virtual_site` | [advertise_custom.advertise_where.virtual_site.virtual_site](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2122203132110010-1231323300323021-0102022013333331-0132300203132330-1211133313232201-1030013232132003-2121333232200202-3313030102012203) |
| `advertise_custom.advertise_where.virtual_site.virtual_site.name` | [advertise_custom.advertise_where.virtual_site.virtual_site.name](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2212112132232010-3110203312111013-2123013202311100-0332231003322100-0122322023121103-0311320132000002-3103232120223132-3110012320313003) |
| `advertise_custom.advertise_where.virtual_site.virtual_site.namespace` | [advertise_custom.advertise_where.virtual_site.virtual_site.namespace](data-sources--udp_loadbalancer--reference--group-001.md#canonical-3032232022301220-1213230032330330-1111301331220032-0303203132111202-3202021310010302-1320212030332203-3011011031221220-0320110231321300) |
| `advertise_custom.advertise_where.virtual_site.virtual_site.tenant` | [advertise_custom.advertise_where.virtual_site.virtual_site.tenant](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2132021313230221-2311301123320230-2232003031013032-3132310000033112-2223210100133013-0003110122230202-1122123302023221-1233100223113113) |
| `advertise_custom.advertise_where.virtual_site_with_vip` | [advertise_custom.advertise_where.virtual_site_with_vip](data-sources--udp_loadbalancer--reference--group-001.md#canonical-0330110333112131-1121203013212202-2330230233030330-3111220113210323-3320201312003312-0210331330331031-1012210002313213-2232223320113002) |
| `advertise_custom.advertise_where.virtual_site_with_vip.ip` | [advertise_custom.advertise_where.virtual_site_with_vip.ip](data-sources--udp_loadbalancer--reference--group-001.md#canonical-3332333313331103-0322312213002022-0121220102123300-2211321202313331-1201300233112332-2311001032312103-2322222012023100-2020332133312303) |
| `advertise_custom.advertise_where.virtual_site_with_vip.network` | [advertise_custom.advertise_where.virtual_site_with_vip.network](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1232111333201111-1031132020003112-0202223111033313-0211120220323031-0213002323223031-1130023120200321-1223310020213122-3132003000231000) |
| `advertise_custom.advertise_where.virtual_site_with_vip.virtual_site` | [advertise_custom.advertise_where.virtual_site_with_vip.virtual_site](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1233022300012223-0303033303220000-3320010233033320-0002310033022002-3130032311102100-1203331232300120-2223132213111102-0232031110302101) |
| `advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.name` | [advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.name](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2022313133110032-1020300321310110-0023010212200233-1023320021010032-0330132003302231-1302210030123002-0211022212030202-3212112210310211) |
| `advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.namespace` | [advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.namespace](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2230200232211301-3111120110132202-1311011231212212-1300213020203013-0200221001112113-2111001301121132-2100332112111121-0032213311012323) |
| `advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.tenant` | [advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.tenant](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1111120102123103-1020332320103333-2012202323121201-2111322113212023-0112220030131323-2032131003231132-1232231313202002-3301213122121311) |
| `advertise_custom.advertise_where.vk8s_service` | [advertise_custom.advertise_where.vk8s_service](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2102013110032202-3001232200131133-3300213031202022-3003323113120331-3130132231012201-2230100031023001-2130113211011313-3311230111330001) |
| `advertise_custom.advertise_where.vk8s_service.site` | [advertise_custom.advertise_where.vk8s_service.site](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1131023232332023-1103321302120300-2012211012212221-1211230312030203-3313220231211331-3302220101221010-1023002101312012-2323220020031020) |
| `advertise_custom.advertise_where.vk8s_service.site.name` | [advertise_custom.advertise_where.vk8s_service.site.name](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2303003122203213-2010303320113203-1233323023130222-0031212113313313-1323132333003131-0332113030020131-2221232211113301-1210212230230121) |
| `advertise_custom.advertise_where.vk8s_service.site.namespace` | [advertise_custom.advertise_where.vk8s_service.site.namespace](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1300102000121233-1311123331032301-0020211330113022-0022230323133032-1001211111303102-3133120132311023-3021203313000022-1001230303123311) |
| `advertise_custom.advertise_where.vk8s_service.site.tenant` | [advertise_custom.advertise_where.vk8s_service.site.tenant](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2010223221021223-0200000210113102-0131230232311031-3310131312101301-1231113202031031-1220110310322113-1031203122233031-1232310030111132) |
| `advertise_custom.advertise_where.vk8s_service.virtual_site` | [advertise_custom.advertise_where.vk8s_service.virtual_site](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1331233233122113-2010013322120313-0103030312013322-2031330322031111-3231022203002300-1023003020222213-1303300031131111-2013300100220130) |
| `advertise_custom.advertise_where.vk8s_service.virtual_site.name` | [advertise_custom.advertise_where.vk8s_service.virtual_site.name](data-sources--udp_loadbalancer--reference--group-001.md#canonical-0013102222333331-0230101033213033-1022130221212330-2002113333321102-0110301310331110-3321333130220233-2003111113012103-2033323212231202) |
| `advertise_custom.advertise_where.vk8s_service.virtual_site.namespace` | [advertise_custom.advertise_where.vk8s_service.virtual_site.namespace](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1000333020313312-0210101210101211-0323333102103233-1013201122333331-1222311311202313-0233311212130013-3323003312103221-0222333122012101) |
| `advertise_custom.advertise_where.vk8s_service.virtual_site.tenant` | [advertise_custom.advertise_where.vk8s_service.virtual_site.tenant](data-sources--udp_loadbalancer--reference--group-001.md#canonical-3221123231121320-1321201112110111-1002310333302002-3133210100001233-1333311000010232-1203010022002313-1123021033323213-2100202311103313) |
| `advertise_on_public` | [advertise_on_public](data-sources--udp_loadbalancer--reference--group-001.md#canonical-3002102112331221-0230030210203322-3322133332201112-0022012021112003-1203203001313113-2100323231233212-2302013110132131-1303033330030333) |
| `advertise_on_public.public_ip` | [advertise_on_public.public_ip](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2023311310110331-1312203213301112-3102120203010222-2220111313321003-3202213230233321-0121232002122300-2111312131330110-3110130030033112) |
| `advertise_on_public.public_ip.name` | [advertise_on_public.public_ip.name](data-sources--udp_loadbalancer--reference--group-001.md#canonical-3002000012212032-0020030321033103-0223103033031101-1322131032203223-2310333310220303-1211003003310230-1333212100230122-0223100213332121) |
| `advertise_on_public.public_ip.namespace` | [advertise_on_public.public_ip.namespace](data-sources--udp_loadbalancer--reference--group-001.md#canonical-0002312023330102-3132100023211031-3320003130112112-3323012020320011-3013130021221301-0001121231011231-0321021121030312-3320201211102303) |
| `advertise_on_public.public_ip.tenant` | [advertise_on_public.public_ip.tenant](data-sources--udp_loadbalancer--reference--group-001.md#canonical-0130312323000223-1010300012133130-1211333203303222-2032133023130001-2010332333022122-2320312101210212-1300103332210102-1233010332202011) |
| `advertise_on_public_default_vip` | [advertise_on_public_default_vip](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2310321122112123-0123232131321211-3313201213130111-1230200200322030-3023023100331123-3301223220012322-2310311120302030-3020231201110300) |
| `annotations` | [annotations](data-sources--udp_loadbalancer--reference--group-001.md#canonical-3310231033303130-2223321111200202-0023002323121012-1001212300302022-3030301202212131-0231230123212333-3303013211031222-1333202200302321) |
| `description` | [description](data-sources--udp_loadbalancer--reference--group-001.md#canonical-3000313021223332-1220013131010233-2230023010231310-1333333330320333-3302232002332101-2203200120311022-3132102010233010-3103111233310303) |
| `dns_volterra_managed` | [dns_volterra_managed](data-sources--udp_loadbalancer--reference--group-001.md#canonical-3323033033220322-2031333033002211-0331002300030032-0302210300302001-2010323221231102-2222330333003121-3232100103223012-1203203303313200) |
| `do_not_advertise` | [do_not_advertise](data-sources--udp_loadbalancer--reference--group-001.md#canonical-0101121300202333-1021331103221322-2001200113112121-3023123100313000-2133032303112201-3112321013203203-1300133202010110-3233112123211320) |
| `domains` | [domains](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1220132012302321-0320120103002132-2121222003133022-1202210011223300-3200221120002220-2222221230220010-3021101213000131-1311110130020210) |
| `hash_policy_choice_random` | [hash_policy_choice_random](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1122223103333023-1233133133031020-1222333013212203-1211020022330121-0311223010102111-0301000123021322-3200002323120202-1133011121110231) |
| `hash_policy_choice_round_robin` | [hash_policy_choice_round_robin](data-sources--udp_loadbalancer--reference--group-001.md#canonical-3301230302200220-2012010210111020-1011232100133300-2020330312230202-3133231231333212-0110212313312013-3023101332333311-3030201230022130) |
| `hash_policy_choice_source_ip_stickiness` | [hash_policy_choice_source_ip_stickiness](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2222212012222101-3221310030001111-3333212101313002-3210123130010303-1113113103110210-0132231120331121-2220203130230002-0122030111023130) |
| `id` | [ID](data-sources--udp_loadbalancer--reference--group-001.md#canonical-3213311023103102-1313122013313231-1222220121211033-3333001323230220-1010032110332001-0032031102032223-0132232200210310-1230231023112313) |
| `idle_timeout` | [idle_timeout](data-sources--udp_loadbalancer--reference--group-001.md#canonical-0011213002103221-2110330321003022-1031132023330022-3113000222313032-0220200222012330-1133013233331223-3311021203122330-0031132330010100) |
| `labels` | [labels](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1022302202222302-0312132100003212-2333131133200332-0230102122103120-0231323013301303-1311320220033112-1222133310103311-3211312130102232) |
| `listen_port` | [listen_port](data-sources--udp_loadbalancer--reference--group-001.md#canonical-0201211311302131-0033213301132110-3130311012101302-1113311023232122-3230023121311030-1332222300233120-2333303030131113-3223211023313010) |
| `name` | [name](data-sources--udp_loadbalancer--reference--group-001.md#canonical-0000021102112212-2221111231031123-1302031222220131-0011223121301113-0231310113103100-2001103212303121-0232131222233002-1113212102321313) |
| `namespace` | [namespace](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2223113313130230-3120203122300133-2110110322032132-3221102032030230-1102113012202220-2012113232122330-0223301321110130-3223212302012332) |
| `no_service_policies` | [no_service_policies](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1032213120323321-0333213321123321-0302320330333313-0032020301112220-3013111302213210-3231332120102320-2113100110303220-1220112311233030) |
| `origin_pools_weights` | [origin_pools_weights](data-sources--udp_loadbalancer--reference--group-001.md#canonical-0133011002012002-1223131320102012-0223032003033302-3120000021210220-3211030220302010-1303101301123100-2003300312220212-2222230333102230) |
| `origin_pools_weights.cluster` | [origin_pools_weights.cluster](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1031211211320302-0031131212201002-0213103313332201-0310321132230101-2230211130210321-3311132102223222-2211022320312032-0203230021120210) |
| `origin_pools_weights.cluster.name` | [origin_pools_weights.cluster.name](data-sources--udp_loadbalancer--reference--group-001.md#canonical-3231203300033101-0101013201221012-3020103010033332-1200322312113010-3010313223001331-1033311313102112-0312120013210303-3211113133221023) |
| `origin_pools_weights.cluster.namespace` | [origin_pools_weights.cluster.namespace](data-sources--udp_loadbalancer--reference--group-001.md#canonical-3230222231000021-2130103220030303-3223213321211322-2001132112020121-0121222113103200-1311313113131100-1102123322321312-3031221311222211) |
| `origin_pools_weights.cluster.tenant` | [origin_pools_weights.cluster.tenant](data-sources--udp_loadbalancer--reference--group-001.md#canonical-3030322101121221-2110220133302203-0222212302322330-1000121101200122-2200303231203313-2233233102302022-0112123022030133-1120132332102320) |
| `origin_pools_weights.endpoint_subsets` | [origin_pools_weights.endpoint_subsets](data-sources--udp_loadbalancer--reference--group-001.md#canonical-0110132010111023-1201101220213321-0112302210211033-0112013012312031-3031003010202233-3023232011231300-0333230011331310-2213332030230103) |
| `origin_pools_weights.pool` | [origin_pools_weights.pool](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1323133101200322-1213302302201112-1032111111213103-3301221020002102-1323332123333211-1303212200321110-2330200321113122-2230233312313103) |
| `origin_pools_weights.pool.name` | [origin_pools_weights.pool.name](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1230102213020221-3212202003330320-0102302113020000-1313031213130322-0311213030333120-2020110332131032-3012020223102113-0320020303212000) |
| `origin_pools_weights.pool.namespace` | [origin_pools_weights.pool.namespace](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2321200220223222-3332100322023130-1031031332012302-3330220231301033-3302103223011033-1212121010332100-1033323111013100-2320003012001213) |
| `origin_pools_weights.pool.tenant` | [origin_pools_weights.pool.tenant](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1031331003021331-0102101320131111-0220220322021111-1212121010311210-2111000303222100-3032011112030313-2020231112122331-3113013103110130) |
| `origin_pools_weights.priority` | [origin_pools_weights.priority](data-sources--udp_loadbalancer--reference--group-001.md#canonical-0300133201021131-0102133022313220-1232323133310202-3210130133113131-1110032212101112-1122320120332321-1313023210331221-3200003102111002) |
| `origin_pools_weights.weight` | [origin_pools_weights.weight](data-sources--udp_loadbalancer--reference--group-001.md#canonical-3021320002112111-3000220113303332-1010323023223120-2113122100112001-0021323333213223-2112110100301011-0111002113210303-2122311220331223) |
| `port_ranges` | [port_ranges](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2032300322031010-2200011221101230-3213232331000322-1032022312031312-2022122301233121-3000300302311221-3322033023022132-3103220223211103) |
| `service_policies_from_namespace` | [service_policies_from_namespace](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1031233021121231-2321310131030302-1012021303323303-3323220310330020-1312131033213310-1332222212230013-3001022313101010-0030103232201320) |
| `udp` | [udp](data-sources--udp_loadbalancer--reference--group-001.md#canonical-0322233120323103-3131330232221001-0022310300013103-1231002210112131-3023010133201203-2123111033231330-1330110310311130-3111201100222010) |

<a id="canonical-2002021201322102-1313022001002303-2111213300230233-1332310003331133-2103202231301002-3220131230302030-2023222322300202-1023100330032210"></a>

## Next pages — Property reference / 012202002322 / 16

- [active_service_policies](data-sources--udp_loadbalancer--reference--group-001.md#canonical-3012132102203030-3133031323022332-1202102030001302-2101121232121222-1230212330110021-3133230213123232-3113221120113321-3020002310301231)
- [advertise_custom](data-sources--udp_loadbalancer--reference--group-001.md#canonical-0102011322232323-2032232030221030-3013012213003033-1131020131301123-2120010001223303-0101100312231222-3210322122303222-2321103220212112)
- [advertise_on_public](data-sources--udp_loadbalancer--reference--group-001.md#canonical-3111011220331032-2313033231112301-0320230202101033-3101213131331212-0221323210201200-3230133203131100-3011303013112000-3303002032321131)
- [advertise_on_public_default_vip](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2311322310303223-0231020022011110-0310223232001233-0311013120223030-0200302330003100-2330123303103011-3100022130033000-2112131213103102)
- [do_not_advertise](data-sources--udp_loadbalancer--reference--group-001.md#canonical-0131000331300322-0103301113212111-2211110212313310-3002002230210313-3103111021003321-1003003123210201-3122100201032330-0222313230310031)
- [hash_policy_choice_random](data-sources--udp_loadbalancer--reference--group-001.md#canonical-3002201000313203-2123320202312030-0022330303110220-0230302222121120-3201231231321301-0232331300203223-0022233321021130-0002331230323021)
- [hash_policy_choice_round_robin](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1223232202021103-2020233233211000-0201322313330131-2301121001023022-2021122103113111-2320302310333212-3013001031012230-2333021133101031)
- [hash_policy_choice_source_ip_stickiness](data-sources--udp_loadbalancer--reference--group-001.md#canonical-0312203212031200-0300032003330222-3023010030101230-1213021131301013-0011220230203121-3021031233331130-0320012111103302-1133302313101320)
- [no_service_policies](data-sources--udp_loadbalancer--reference--group-001.md#canonical-0311102003233330-0121133300020322-3132013220222322-3203312022330021-3033322310112312-2323110002200030-0120301333121123-3011112033313013)
- [origin_pools_weights](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2320311033331002-1133320121212213-0320333113110131-2120100213202221-2233020011030301-1322013030212121-1131120221033210-2032330100122213)
- [service_policies_from_namespace](data-sources--udp_loadbalancer--reference--group-001.md#canonical-3122231333000123-0312111333033323-2323101223110320-3223113220211321-0222023230110012-1103201221010012-0210212333003320-3122113112330300)
- [udp](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2320000221111121-2211321333322221-0131123101031120-1011222001302102-3020022130323223-1330300011100301-1023000101320102-3123222213330010)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)

<a id="canonical-3012132102203030-3133031323022332-1202102030001302-2101121232121222-1230212330110021-3133230213123232-3113221120113321-3020002310301231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2323332110010021-2311203321011111-0321132131102113-2323332303130123-2003332301121120-1233311321131230-1132113303221122-0021023330020020"></a>

## active_service_policies — active_service_policies / 130113211111 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2100233330113013-1011020101321013-1113020011303120-3332020200220223-1130023001010003-0211031020201213-0121011023231220-3020101330032130)
- active_service_policies

<a id="canonical-3022332123102210-1013032320232232-1231133223023010-0131231223312023-0230212023013212-1001032313131122-2233022323231012-2120301003323332"></a>

Type: `"single"`. Computed.

\[OneOf: active\_service\_policies, no\_service\_policies, service\_policies\_from\_namespace;
Default: no\_service\_policies\] Configuration parameter for active service policies.

Upstream description:

List of service policies.

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

OneOf alternatives in this subsection:

- [active_service_policies](data-sources--udp_loadbalancer--reference--group-001.md#canonical-3022332123102210-1013032320232232-1231133223023010-0131231223312023-0230212023013212-1001032313131122-2233022323231012-2120301003323332)
- [no_service_policies](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1032213120323321-0333213321123321-0302320330333313-0032020301112220-3013111302213210-3231332120102320-2113100110303220-1220112311233030)
- [service_policies_from_namespace](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1031233021121231-2321310131030302-1012021303323303-3323220310330020-1312131033213310-1332222212230013-3001022313101010-0030103232201320)

Select alternatives according to the provider validators above.

<a id="canonical-2001020333230121-1211313210312100-0001210202021312-0322213033131103-2222132320123112-1203121323320000-3000132030300203-3333312002020302"></a>

## Direct properties — active_service_policies / 130113211111 / 3

- [policies](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1021102200332331-1020232301222201-1013202003001312-3011123313230301-0023101301230022-0032321300032311-2122311100101213-2012011303021203): complete subsection reference.

<a id="canonical-1020010312200021-1033011131102003-3202130331101333-1121201303201223-3030103003131330-0212003213212022-1023033221031321-0221233031022100"></a>

## Next pages — active_service_policies / 130113211111 / 4

- [active_service_policies.policies](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1021102200332331-1020232301222201-1013202003001312-3011123313230301-0023101301230022-0032321300032311-2122311100101213-2012011303021203)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2100233330113013-1011020101321013-1113020011303120-3332020200220223-1130023001010003-0211031020201213-0121011023231220-3020101330032130)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)

<a id="canonical-1021102200332331-1020232301222201-1013202003001312-3011123313230301-0023101301230022-0032321300032311-2122311100101213-2012011303021203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3103011331003221-3121012323023100-3300201312110003-0221200111110330-1020300020323322-0120132311332013-2022013110322323-3311001330213333"></a>

## active_service_policies.policies — policies / 002333002120 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2100233330113013-1011020101321013-1113020011303120-3332020200220223-1130023001010003-0211031020201213-0121011023231220-3020101330032130)
- [active_service_policies](data-sources--udp_loadbalancer--reference--group-001.md#canonical-3012132102203030-3133031323022332-1202102030001302-2101121232121222-1230212330110021-3133230213123232-3113221120113321-3020002310301231)
- active_service_policies.policies

<a id="canonical-2111200323110000-0011021030120010-1031201133001021-1001332130331211-2233120303330032-1312212330030001-2332131213113232-3202002312301120"></a>

Type: `"list"`. Computed.

Service Policies is a sequential engine where policies (and rules within the policy) are evaluated
one after the other. It's important to define the correct order (policies evaluated from top to
bottom in the list) for service policies, to GET the intended result. For each request, its..

Upstream description:

Service Policies is a sequential engine where policies (and rules within the policy) are evaluated
one after the other. It's important to define the correct order (policies evaluated from top to
bottom in the list) for service policies, to GET the intended result. For each request, its
characteristics are evaluated based on the match criteria in each service policy starting at the
top. If there is a match in the current policy, then the policy takes effect, and no more policies
are evaluated. Otherwise, the next policy is evaluated. If all policies are evaluated and none
match, then the request will be denied by default.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0003200302302222-2101111333301311-1033112022110330-3103210302010131-3100222301221231-0332330302100123-3000213300032100-1230200121322220"></a>

## Direct properties — policies / 002333002120 / 3

<a id="canonical-3320111100011310-2001033312312100-1212302033220332-3110222200201312-0232202313211113-1231133133233320-1121310100310100-3303200213201130"></a>

<a id="canonical-0210310132231332-2032202001313303-3221123332231222-1210030333103311-1030322322122033-0012130122320113-2033122020020113-1000033002311330"></a>

## name property — policies / 002333002120 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-2032111202021011-0012310030223232-3211020122232320-0012221321231013-0213222210000123-0030102312002122-0313032132313130-3132002312030031"></a>

<a id="canonical-0313021130312102-0032300001223333-3033012301102222-1121320033301312-1230300331231311-1001310232220323-3031331133231300-3230013102032333"></a>

## namespace property — policies / 002333002120 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
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
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3203011200020133-2130212201032113-2023310333111311-1200222220212133-1311301211123020-1013102303102211-2102330211212013-2201010200112100"></a>

<a id="canonical-3110123223323012-1131120333203213-1201230131232110-3333230203223112-3222010233310122-0302033110120111-3211220133112002-3223133303130023"></a>

## tenant property — policies / 002333002120 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-2110032110002313-0211301200001322-2300220203321111-3213210331302213-3222332312220121-0203111012033122-1103103002322322-2321301010023302"></a>

## Next pages — policies / 002333002120 / 7

- [active_service_policies](data-sources--udp_loadbalancer--reference--group-001.md#canonical-3012132102203030-3133031323022332-1202102030001302-2101121232121222-1230212330110021-3133230213123232-3113221120113321-3020002310301231)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)

<a id="canonical-0102011322232323-2032232030221030-3013012213003033-1131020131301123-2120010001223303-0101100312231222-3210322122303222-2321103220212112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1203322012312331-1110031300203023-2312021110103203-1312033033012030-3022131001302023-3110321030221230-3021323033310203-0323201230311301"></a>

## advertise_custom — advertise_custom / 332303230232 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2100233330113013-1011020101321013-1113020011303120-3332020200220223-1130023001010003-0211031020201213-0121011023231220-3020101330032130)
- advertise_custom

<a id="canonical-1100123320322233-0112312301223220-2303322231013321-3113113310222200-2133011023231220-2301331020100210-3313300312303321-2021233211330012"></a>

Type: `"single"`. Computed.

\[OneOf: advertise\_custom, advertise\_on\_public, advertise\_on\_public\_default\_vip,
do\_not\_advertise; Default: advertise\_on\_public\_default\_vip\] Defines a way to advertise a VIP
on specific sites.

Upstream description:

This defines a way to advertise a VIP on specific sites.

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

OneOf alternatives in this subsection:

- [advertise_custom](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1100123320322233-0112312301223220-2303322231013321-3113113310222200-2133011023231220-2301331020100210-3313300312303321-2021233211330012)
- [advertise_on_public](data-sources--udp_loadbalancer--reference--group-001.md#canonical-3002102112331221-0230030210203322-3322133332201112-0022012021112003-1203203001313113-2100323231233212-2302013110132131-1303033330030333)
- [advertise_on_public_default_vip](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2310321122112123-0123232131321211-3313201213130111-1230200200322030-3023023100331123-3301223220012322-2310311120302030-3020231201110300)
- [do_not_advertise](data-sources--udp_loadbalancer--reference--group-001.md#canonical-0101121300202333-1021331103221322-2001200113112121-3023123100313000-2133032303112201-3112321013203203-1300133202010110-3233112123211320)

Select alternatives according to the provider validators above.

<a id="canonical-2033121203003033-3101321322002033-3100211312202120-1233320121312021-2313003102331231-2102232020210213-3000300112132210-3032110302212213"></a>

## Direct properties — advertise_custom / 332303230232 / 3

- [advertise_where](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2330333120302301-0122113032301122-1110222000313131-0313010302030110-0102002133220333-1223323113301222-1013010221110031-2001020101120113): complete subsection reference.

<a id="canonical-1213020211123021-0032332031233103-1210001032121110-3311002020100123-2323211021023130-2032312301320001-0201033222233000-2330123320002132"></a>

## Next pages — advertise_custom / 332303230232 / 4

- [advertise_custom.advertise_where](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2330333120302301-0122113032301122-1110222000313131-0313010302030110-0102002133220333-1223323113301222-1013010221110031-2001020101120113)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2100233330113013-1011020101321013-1113020011303120-3332020200220223-1130023001010003-0211031020201213-0121011023231220-3020101330032130)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)

<a id="canonical-2330333120302301-0122113032301122-1110222000313131-0313010302030110-0102002133220333-1223323113301222-1013010221110031-2001020101120113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2233232323012102-2321230201230013-0000322021032100-3330013331312133-0220212000330300-1131003322312231-3110122020030220-2321222321031330"></a>

## advertise_custom.advertise_where — advertise_where / 132220020233 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2100233330113013-1011020101321013-1113020011303120-3332020200220223-1130023001010003-0211031020201213-0121011023231220-3020101330032130)
- [advertise_custom](data-sources--udp_loadbalancer--reference--group-001.md#canonical-0102011322232323-2032232030221030-3013012213003033-1131020131301123-2120010001223303-0101100312231222-3210322122303222-2321103220212112)
- advertise_custom.advertise_where

<a id="canonical-3322330333233133-1301212300100033-1023010031000021-3021130333311021-2201310300222023-2003210122303202-0211023100120231-0211113033310011"></a>

Type: `"list"`. Computed.

Where should this load balancer be available.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1120312023332133-1023233032303131-3130300330100122-0200312233233132-2220102002120330-2332220120320220-3332333111210123-2300232221013201"></a>

## Direct properties — advertise_where / 132220020233 / 3

- [advertise_dualstack_on_public](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2101013213232031-2001001021331021-2102121000333122-3001130100121200-1200003311131122-1100213023032120-1320032002210003-0330211012121230): complete subsection reference.

- [advertise_on_public](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1220130212233223-2131001310230201-0323022012123110-2001303312030322-2113313110121202-0300113322133221-2103020120133230-1032220102122232): complete subsection reference.

- [advertise_v6_on_public](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2202121311110000-1101032111330033-0133220103022103-2022122331000302-0033023303010310-3231110021033130-0020301313023202-1100111022011131): complete subsection reference.

<a id="canonical-2112202310132202-3002213133311331-3223231330103332-3332132002030031-2321213122311100-3202232200113302-1121113212231113-1322130113311010"></a>

<a id="canonical-3013001130011332-2312032223001010-3311112013220002-2023330113321212-0330131312210012-2320213230030320-2223011011333223-0011321211200022"></a>

## port property — advertise_where / 132220020233 / 4

Type: `"number"`. Computed.

Exclusive with \[port\_ranges use\_default\_port\] Port to Listen.

Upstream description:

Exclusive with \[port\_ranges use\_default\_port\] Port to Listen.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-3323132113031300-3220013131120301-1130030032301301-3113032122232223-3223010312203031-2002100212012311-3213332223023300-3321122210223103"></a>

<a id="canonical-2021013310202312-1001310331330011-0001213030210103-3031132203322230-1333011221133033-1020122220132123-2110332233110113-3111030220113212"></a>

## port_ranges property — advertise_where / 132220020233 / 5

Type: `"string"`. Computed.

Exclusive with \[port use\_default\_port\] A string containing a comma separated list of port
ranges. Each port range consists of a single port or two ports separated by '-'.

Upstream description:

Exclusive with \[port use\_default\_port\] A string containing a comma separated list of port
ranges. Each port range consists of a single port or two ports separated by "-".

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
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
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  }
}
```

- [site](data-sources--udp_loadbalancer--reference--group-001.md#canonical-0001321010020213-0203311323323110-0233120231122002-3331103102133122-0331330110330133-3010302003030201-3012012310300300-0310212103130331): complete subsection reference.

- [use_default_port](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2100122031331002-2003002020311110-1200030312012302-1130213020220111-1103122201220211-0000212121123211-0222320133113302-0200220321012100): complete subsection reference.

- [virtual_network](data-sources--udp_loadbalancer--reference--group-001.md#canonical-3220100322102130-0323221211203111-0221213211203323-1020220322130200-0230320321133323-3220103132032220-0112232113302131-3013003122012013): complete subsection reference.

- [virtual_site](data-sources--udp_loadbalancer--reference--group-001.md#canonical-3203131010231322-1001331213122003-3210202200131011-3130020031032230-3031322232223100-3011212301111021-0133201102201212-1303301012202311): complete subsection reference.

- [virtual_site_with_vip](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2212122103220300-0203010212033312-1231031331121312-1122332300022302-1012303121233033-3302333231031311-1332110132333011-2300133232310310): complete subsection reference.

- [vk8s_service](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1322123110212212-1033222023310221-1021233002213200-2121220132323000-1120320111033002-2012002311123220-0020131200103231-0323132202303320): complete subsection reference.

<a id="canonical-3133312101113303-3033301013020020-3233210201130133-3300321223011200-1221130130332223-2222233131111222-0332331220003311-0100302003103033"></a>

## Next pages — advertise_where / 132220020233 / 6

- [advertise_custom.advertise_where.advertise_dualstack_on_public](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2101013213232031-2001001021331021-2102121000333122-3001130100121200-1200003311131122-1100213023032120-1320032002210003-0330211012121230)
- [advertise_custom.advertise_where.advertise_on_public](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1220130212233223-2131001310230201-0323022012123110-2001303312030322-2113313110121202-0300113322133221-2103020120133230-1032220102122232)
- [advertise_custom.advertise_where.advertise_v6_on_public](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2202121311110000-1101032111330033-0133220103022103-2022122331000302-0033023303010310-3231110021033130-0020301313023202-1100111022011131)
- [advertise_custom.advertise_where.site](data-sources--udp_loadbalancer--reference--group-001.md#canonical-0001321010020213-0203311323323110-0233120231122002-3331103102133122-0331330110330133-3010302003030201-3012012310300300-0310212103130331)
- [advertise_custom.advertise_where.use_default_port](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2100122031331002-2003002020311110-1200030312012302-1130213020220111-1103122201220211-0000212121123211-0222320133113302-0200220321012100)
- [advertise_custom.advertise_where.virtual_network](data-sources--udp_loadbalancer--reference--group-001.md#canonical-3220100322102130-0323221211203111-0221213211203323-1020220322130200-0230320321133323-3220103132032220-0112232113302131-3013003122012013)
- [advertise_custom.advertise_where.virtual_site](data-sources--udp_loadbalancer--reference--group-001.md#canonical-3203131010231322-1001331213122003-3210202200131011-3130020031032230-3031322232223100-3011212301111021-0133201102201212-1303301012202311)
- [advertise_custom.advertise_where.virtual_site_with_vip](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2212122103220300-0203010212033312-1231031331121312-1122332300022302-1012303121233033-3302333231031311-1332110132333011-2300133232310310)
- [advertise_custom.advertise_where.vk8s_service](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1322123110212212-1033222023310221-1021233002213200-2121220132323000-1120320111033002-2012002311123220-0020131200103231-0323132202303320)
- [advertise_custom](data-sources--udp_loadbalancer--reference--group-001.md#canonical-0102011322232323-2032232030221030-3013012213003033-1131020131301123-2120010001223303-0101100312231222-3210322122303222-2321103220212112)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)

<a id="canonical-2101013213232031-2001001021331021-2102121000333122-3001130100121200-1200003311131122-1100213023032120-1320032002210003-0330211012121230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1120202321312210-0322010323330213-1222202312133211-2033212012102313-0011101000102121-0121010001303320-2300111031233333-0101233210230312"></a>

## advertise_custom.advertise_where.advertise_dualstack_on_public — advertise_dualstack_on_public / 103233323030 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2100233330113013-1011020101321013-1113020011303120-3332020200220223-1130023001010003-0211031020201213-0121011023231220-3020101330032130)
- [advertise_custom](data-sources--udp_loadbalancer--reference--group-001.md#canonical-0102011322232323-2032232030221030-3013012213003033-1131020131301123-2120010001223303-0101100312231222-3210322122303222-2321103220212112)
- [advertise_custom.advertise_where](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2330333120302301-0122113032301122-1110222000313131-0313010302030110-0102002133220333-1223323113301222-1013010221110031-2001020101120113)
- advertise_custom.advertise_where.advertise_dualstack_on_public

<a id="canonical-1030321213303022-1003212332131122-1303000033113112-2023313123223101-3311302103001210-3112102003131213-1221202133220100-3331002233000103"></a>

Type: `"single"`. Computed.

Defines a way to advertise a load balancer on public. If optional public\_ip is provided, it will
only be advertised on RE sites where that public\_ip is available.

Upstream description:

This defines a way to advertise a load balancer on public. If optional public\_ip is provided, it
will only be advertised on RE sites where that public\_ip is available.

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

<a id="canonical-2310322121012321-0201212122001320-3010111031130320-0031332222003101-2233022121111111-0210332011022202-0320011232111310-0113120203202232"></a>

## Direct properties — advertise_dualstack_on_public / 103233323030 / 3

- [public_ip](data-sources--udp_loadbalancer--reference--group-001.md#canonical-0111022201132312-2133003322313033-0231313001221012-3331300321203121-3322200132120313-0332130300202013-3013220332322331-1033231222300222): complete subsection reference.

<a id="canonical-2200030302333312-3330330122002320-0220031100231122-2030011131031130-0030211123301021-3221003303131102-3100012221300020-1212120011133232"></a>

## Next pages — advertise_dualstack_on_public / 103233323030 / 4

- [advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip](data-sources--udp_loadbalancer--reference--group-001.md#canonical-0111022201132312-2133003322313033-0231313001221012-3331300321203121-3322200132120313-0332130300202013-3013220332322331-1033231222300222)
- [advertise_custom.advertise_where](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2330333120302301-0122113032301122-1110222000313131-0313010302030110-0102002133220333-1223323113301222-1013010221110031-2001020101120113)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)

<a id="canonical-0111022201132312-2133003322313033-0231313001221012-3331300321203121-3322200132120313-0332130300202013-3013220332322331-1033231222300222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1003133323012032-0130103302121102-3200030030030013-3312012033320231-0032223311023102-2111310202101321-1003233100301212-3302230010201312"></a>

## advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip — public_ip / 121330112121 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2100233330113013-1011020101321013-1113020011303120-3332020200220223-1130023001010003-0211031020201213-0121011023231220-3020101330032130)
- [advertise_custom](data-sources--udp_loadbalancer--reference--group-001.md#canonical-0102011322232323-2032232030221030-3013012213003033-1131020131301123-2120010001223303-0101100312231222-3210322122303222-2321103220212112)
- [advertise_custom.advertise_where](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2330333120302301-0122113032301122-1110222000313131-0313010302030110-0102002133220333-1223323113301222-1013010221110031-2001020101120113)
- [advertise_custom.advertise_where.advertise_dualstack_on_public](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2101013213232031-2001001021331021-2102121000333122-3001130100121200-1200003311131122-1100213023032120-1320032002210003-0330211012121230)
- advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip

<a id="canonical-1130020210100000-3220101201001003-0122223222111120-0232122103001221-2130212311233020-1001011112012322-3311120233321203-1230031032203212"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-3113133122213303-1301002023212230-0330233012010111-1303232330101131-2032023211133111-2211303020232103-0301011000010233-3013133202000100"></a>

## Direct properties — public_ip / 121330112121 / 3

<a id="canonical-2321202101022130-1201032203211122-2001222101232303-1130213332232000-1100013333003100-3013110232231002-2111320012212131-1121120033213312"></a>

<a id="canonical-0101210122130301-2010023213232201-2030312023223030-1303223232102022-1021121023231313-0110013231302103-0322030311103121-3123110211110130"></a>

## name property — public_ip / 121330112121 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-1020023233112120-1221312003112000-3003230211120323-1112330332202333-0022113211233223-0032021323120312-3113200030230001-3122233121021323"></a>

<a id="canonical-3012321002132122-2010211311230102-3232100323223113-2230300321112031-2110110230013332-2313102132331121-1030030323030033-2223022123310213"></a>

## namespace property — public_ip / 121330112121 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
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
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-2213223330003300-2311332312310010-0330031023012313-1233000121231120-3331222012313133-2212112212002110-1022123131013121-2220013232133100"></a>

<a id="canonical-1111300203130211-2120302021232032-1330022321113311-0303222101103001-1031322211021313-3230012002302120-3213330021122303-1110010002201102"></a>

## tenant property — public_ip / 121330112121 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-0301330203231211-3112312032312232-3233030120011113-0330000121210003-1200121330333221-0012022021110212-2130231222022232-1001300332302132"></a>

## Next pages — public_ip / 121330112121 / 7

- [advertise_custom.advertise_where.advertise_dualstack_on_public](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2101013213232031-2001001021331021-2102121000333122-3001130100121200-1200003311131122-1100213023032120-1320032002210003-0330211012121230)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)

<a id="canonical-1220130212233223-2131001310230201-0323022012123110-2001303312030322-2113313110121202-0300113322133221-2103020120133230-1032220102122232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1332322002023311-3201232322022303-1220131302230120-1321321213103303-0022322312213231-2123100232320203-3231020223032130-0222200001021221"></a>

## advertise_custom.advertise_where.advertise_on_public — advertise_on_public / 120203212312 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2100233330113013-1011020101321013-1113020011303120-3332020200220223-1130023001010003-0211031020201213-0121011023231220-3020101330032130)
- [advertise_custom](data-sources--udp_loadbalancer--reference--group-001.md#canonical-0102011322232323-2032232030221030-3013012213003033-1131020131301123-2120010001223303-0101100312231222-3210322122303222-2321103220212112)
- [advertise_custom.advertise_where](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2330333120302301-0122113032301122-1110222000313131-0313010302030110-0102002133220333-1223323113301222-1013010221110031-2001020101120113)
- advertise_custom.advertise_where.advertise_on_public

<a id="canonical-0033320301023120-3000102012021310-3111331133303311-0022231210213200-3201110032310220-2132032110330313-3210013312332232-0112032122020013"></a>

Type: `"single"`. Computed.

Defines a way to advertise a load balancer on public. If optional public\_ip is provided, it will
only be advertised on RE sites where that public\_ip is available.

Upstream description:

This defines a way to advertise a load balancer on public. If optional public\_ip is provided, it
will only be advertised on RE sites where that public\_ip is available.

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

<a id="canonical-3310330322331223-2011113031333331-3023001123212331-3001312321022131-2110023231110322-1011131131212212-0303201112123313-3220203010330112"></a>

## Direct properties — advertise_on_public / 120203212312 / 3

- [public_ip](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2011210233013212-2023012310131023-2133230312111031-2132203031032213-3311020220010010-1100310003000313-3332213111201030-3312231131113033): complete subsection reference.

<a id="canonical-2023031003122111-1113101012310323-2031320313132100-2210003011033322-2332321112211310-0233322003232212-0200313202121220-3221220002203212"></a>

## Next pages — advertise_on_public / 120203212312 / 4

- [advertise_custom.advertise_where.advertise_on_public.public_ip](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2011210233013212-2023012310131023-2133230312111031-2132203031032213-3311020220010010-1100310003000313-3332213111201030-3312231131113033)
- [advertise_custom.advertise_where](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2330333120302301-0122113032301122-1110222000313131-0313010302030110-0102002133220333-1223323113301222-1013010221110031-2001020101120113)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)

<a id="canonical-2011210233013212-2023012310131023-2133230312111031-2132203031032213-3311020220010010-1100310003000313-3332213111201030-3312231131113033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3223223310212000-3102022210100212-3310000112223201-0233331332331101-1133002301231023-1221303312232021-1331100322321020-2021221022331233"></a>

## advertise_custom.advertise_where.advertise_on_public.public_ip — public_ip / 003203320000 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2100233330113013-1011020101321013-1113020011303120-3332020200220223-1130023001010003-0211031020201213-0121011023231220-3020101330032130)
- [advertise_custom](data-sources--udp_loadbalancer--reference--group-001.md#canonical-0102011322232323-2032232030221030-3013012213003033-1131020131301123-2120010001223303-0101100312231222-3210322122303222-2321103220212112)
- [advertise_custom.advertise_where](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2330333120302301-0122113032301122-1110222000313131-0313010302030110-0102002133220333-1223323113301222-1013010221110031-2001020101120113)
- [advertise_custom.advertise_where.advertise_on_public](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1220130212233223-2131001310230201-0323022012123110-2001303312030322-2113313110121202-0300113322133221-2103020120133230-1032220102122232)
- advertise_custom.advertise_where.advertise_on_public.public_ip

<a id="canonical-2110111123000211-1331330211201330-2202303103310300-0032001002333122-1031322000330101-1011221001231033-3103212332022131-3033120323233320"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-2333111313223000-0312020123112231-1112130130331013-1230020333020111-2303003031001020-0320102232111323-0102122311001303-1133321323002223"></a>

## Direct properties — public_ip / 003203320000 / 3

<a id="canonical-0323022032003232-2123332123201120-2112233122213112-1232013333102333-0100133130033203-1311122032130012-1120330011130022-3211330010333010"></a>

<a id="canonical-0013031313202112-1001201121111002-0132231223131020-3132220213021313-0013103112110311-2200300303021202-3022133123123330-3100310210033031"></a>

## name property — public_ip / 003203320000 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-3131023030100001-0130221113200100-0302212220013230-2202023022301012-3231321212213013-2220320211323030-0333012230010123-1230220232131112"></a>

<a id="canonical-2333131031232010-0121300022013010-2313311233001013-0310112333000212-0211120121300321-1230111110130020-1313321303133020-3302302123222202"></a>

## namespace property — public_ip / 003203320000 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
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
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-0321300331003202-0320332303121011-0210011021002031-2010302301301203-3100031010321003-1012130231323230-0313113321121003-1310230201021112"></a>

<a id="canonical-3120313211022131-3131120320101000-1200112132200211-2002023023220223-3210300102033221-2201210111012022-3132231202310221-1000130022001201"></a>

## tenant property — public_ip / 003203320000 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-0330231111213021-0023003221102203-1213033212021303-3313331302311310-0033223332131223-0032022111303333-3133123012020200-3203123213300322"></a>

## Next pages — public_ip / 003203320000 / 7

- [advertise_custom.advertise_where.advertise_on_public](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1220130212233223-2131001310230201-0323022012123110-2001303312030322-2113313110121202-0300113322133221-2103020120133230-1032220102122232)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)

<a id="canonical-2202121311110000-1101032111330033-0133220103022103-2022122331000302-0033023303010310-3231110021033130-0020301313023202-1100111022011131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1100310323203131-0212232132332023-2223223021001222-3100111000212221-3130103220230121-1311021100222022-1223330200303321-2331020323000230"></a>

## advertise_custom.advertise_where.advertise_v6_on_public — advertise_v6_on_public / 300111233213 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2100233330113013-1011020101321013-1113020011303120-3332020200220223-1130023001010003-0211031020201213-0121011023231220-3020101330032130)
- [advertise_custom](data-sources--udp_loadbalancer--reference--group-001.md#canonical-0102011322232323-2032232030221030-3013012213003033-1131020131301123-2120010001223303-0101100312231222-3210322122303222-2321103220212112)
- [advertise_custom.advertise_where](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2330333120302301-0122113032301122-1110222000313131-0313010302030110-0102002133220333-1223323113301222-1013010221110031-2001020101120113)
- advertise_custom.advertise_where.advertise_v6_on_public

<a id="canonical-2220113100301322-2223200100012202-3110103322221221-1010221000133002-0333132110123033-3121023021102303-2302033010223311-2003302310013332"></a>

Type: `"single"`. Computed.

Defines a way to advertise a load balancer on public. If optional public\_ip is provided, it will
only be advertised on RE sites where that public\_ip is available.

Upstream description:

This defines a way to advertise a load balancer on public. If optional public\_ip is provided, it
will only be advertised on RE sites where that public\_ip is available.

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

<a id="canonical-1031000211210130-2230013212311303-1221013023321322-1020331032211101-2021212300231323-2020310111122000-2323110011121103-3021313111333022"></a>

## Direct properties — advertise_v6_on_public / 300111233213 / 3

- [public_ip](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1210121301330131-0322021330222021-3333113213011213-0223230320212011-3313223201230101-1301203222230002-3023303120202312-2023022013303012): complete subsection reference.

<a id="canonical-0303110033120030-0222113213323031-3133031011100333-0102221122001332-0311220130120230-2332211112033232-2332023101210132-1300020312030100"></a>

## Next pages — advertise_v6_on_public / 300111233213 / 4

- [advertise_custom.advertise_where.advertise_v6_on_public.public_ip](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1210121301330131-0322021330222021-3333113213011213-0223230320212011-3313223201230101-1301203222230002-3023303120202312-2023022013303012)
- [advertise_custom.advertise_where](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2330333120302301-0122113032301122-1110222000313131-0313010302030110-0102002133220333-1223323113301222-1013010221110031-2001020101120113)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)

<a id="canonical-1210121301330131-0322021330222021-3333113213011213-0223230320212011-3313223201230101-1301203222230002-3023303120202312-2023022013303012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0221322301112111-0101113313331023-1001003200001320-1302210232013200-3303030233200200-2112200312222201-3302213301031331-0100320213131131"></a>

## advertise_custom.advertise_where.advertise_v6_on_public.public_ip — public_ip / 303203300022 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2100233330113013-1011020101321013-1113020011303120-3332020200220223-1130023001010003-0211031020201213-0121011023231220-3020101330032130)
- [advertise_custom](data-sources--udp_loadbalancer--reference--group-001.md#canonical-0102011322232323-2032232030221030-3013012213003033-1131020131301123-2120010001223303-0101100312231222-3210322122303222-2321103220212112)
- [advertise_custom.advertise_where](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2330333120302301-0122113032301122-1110222000313131-0313010302030110-0102002133220333-1223323113301222-1013010221110031-2001020101120113)
- [advertise_custom.advertise_where.advertise_v6_on_public](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2202121311110000-1101032111330033-0133220103022103-2022122331000302-0033023303010310-3231110021033130-0020301313023202-1100111022011131)
- advertise_custom.advertise_where.advertise_v6_on_public.public_ip

<a id="canonical-3020231130121203-0003233310222302-0020013031023121-2201100112213010-0102012330310221-0031030020033103-1302030302333303-2323210130303201"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-1010011311101111-2320120330221021-3103112310022032-0021332003133121-2123331031301212-1302210131311111-0123100032001302-2221022300033202"></a>

## Direct properties — public_ip / 303203300022 / 3

<a id="canonical-1301123013232111-0323022112113331-2020023302320321-1100112000102201-1013001110313010-3300322301100303-2302101232112232-3312123010232131"></a>

<a id="canonical-2203332012230031-1102021013102312-0201301230001313-3013302310211300-2310000113011130-2203210222210202-3132000121311103-3313223120220031"></a>

## name property — public_ip / 303203300022 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-1103210133320333-3003010213100021-0300223211021323-1212210302330230-3200213002331213-0300301220032332-0033203300031030-1101120301000011"></a>

<a id="canonical-3110023013203331-0001003001231122-3232201102203201-1322021111102333-0233123330230312-0111112210330111-0022331120021003-0200003222021220"></a>

## namespace property — public_ip / 303203300022 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
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
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3013223323001300-3220120133021032-0210333132020301-1032331332122020-3001312231102310-0311321030303233-0333222131220002-1000101002020200"></a>

<a id="canonical-1133303213101130-1122130211302200-2122310231232123-2200322102103010-3210102123131003-0313331132333103-3221201333121103-2302010020211100"></a>

## tenant property — public_ip / 303203300022 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-0013111022211131-3223032200231332-3211032232202320-1213103200333002-3013201032330302-0331131301303012-0112022100122010-3110100000231111"></a>

## Next pages — public_ip / 303203300022 / 7

- [advertise_custom.advertise_where.advertise_v6_on_public](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2202121311110000-1101032111330033-0133220103022103-2022122331000302-0033023303010310-3231110021033130-0020301313023202-1100111022011131)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)

<a id="canonical-0001321010020213-0203311323323110-0233120231122002-3331103102133122-0331330110330133-3010302003030201-3012012310300300-0310212103130331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2122231122031012-3311111011321222-3102200322013022-2221201301101021-0332123032022012-0030010120101232-3333031123331130-3011202113131220"></a>

## advertise_custom.advertise_where.site — site / 030122211101 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2100233330113013-1011020101321013-1113020011303120-3332020200220223-1130023001010003-0211031020201213-0121011023231220-3020101330032130)
- [advertise_custom](data-sources--udp_loadbalancer--reference--group-001.md#canonical-0102011322232323-2032232030221030-3013012213003033-1131020131301123-2120010001223303-0101100312231222-3210322122303222-2321103220212112)
- [advertise_custom.advertise_where](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2330333120302301-0122113032301122-1110222000313131-0313010302030110-0102002133220333-1223323113301222-1013010221110031-2001020101120113)
- advertise_custom.advertise_where.site

<a id="canonical-2212221312130111-0232110003130221-1010331231320201-2200321311120002-3120000021232100-2120130010132133-2222011333033003-2021210321221311"></a>

Type: `"single"`. Computed.

Defines a reference to a CE site along with network type and an optional IP address where a load
balancer could be advertised.

Upstream description:

This defines a reference to a CE site along with network type and an optional IP address where a
load balancer could be advertised.

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

<a id="canonical-1031202120230300-3320123301222032-0300321231321013-3302231110212312-1323111102130030-3230033023033023-2332303021123330-0032122320320002"></a>

## Direct properties — site / 030122211101 / 3

<a id="canonical-0111013013022111-0000002233013230-3311100111002020-1103111001021332-2233101123120301-0002112010220303-3321213311210321-3330312022113111"></a>

<a id="canonical-1233023031322203-3100122231312221-3313331120201223-3032000313201002-0223233020002213-3210103001010322-0222020020333213-1230033011232103"></a>

## ip property — site / 030122211101 / 4

Type: `"string"`. Computed.

Use given IP address as VIP on the site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-0000121300333031-3121213202231132-1100102210022223-0330301302000013-3101322220020231-0002020222303122-2112312032021233-1110101312012100"></a>

<a id="canonical-3313332303103012-2102023331303020-2212223213001011-2023000023201300-1000003213222231-3011131321200331-2322322100213210-1330230121120211"></a>

## network property — site / 030122211101 / 5

Type: `"string"`. Computed.

\[Enum:
SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE|SITE\_NETWORK\_INSIDE|SITE\_NETWORK\_OUTSIDE|SITE\_NETWORK\_SERVICE|SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_IP\_FABRIC\]
Defines network types to be used on site All inside and outside networks. All inside and outside
networks with internet VIP support. All inside networks. Possible values are
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`, \`SITE\_NETWORK\_INSIDE\`, \`SITE\_NETWORK\_OUTSIDE\`,
\`SITE\_NETWORK\_SERVICE\`, \`SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP\`,
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP\`, \`SITE\_NETWORK\_IP\_FABRIC\`.
Defaults to \`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`.

Upstream description:

This defines network types to be used on site

All inside and outside networks. All inside and outside networks with internet VIP support. All
inside networks. All outside networks. All outside networks with internet VIP support. VK8s service
network. &#8203;- SITE\_NETWORK\_IP\_FABRIC: VER IP Fabric network for the site

This Virtual network type is used for exposing virtual host on IP Fabric network on the VER site or
for endpoint in IP Fabric network.

Receipt-pinned upstream constraints:

```json
{
  "default": "SITE_NETWORK_INSIDE_AND_OUTSIDE",
  "enum": [
    "SITE_NETWORK_INSIDE_AND_OUTSIDE",
    "SITE_NETWORK_INSIDE",
    "SITE_NETWORK_OUTSIDE",
    "SITE_NETWORK_SERVICE",
    "SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_IP_FABRIC"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [site](data-sources--udp_loadbalancer--reference--group-001.md#canonical-0132212221220101-0133210003132200-0222131122221002-0112021001222330-2230231331213111-1230020000100322-1103313030020131-3013030130203221): complete subsection reference.

<a id="canonical-2121330233103033-2033300122013102-0130231011010021-3301032133120021-3222310023121303-1033223232321023-0011301313013222-0202103020101132"></a>

## Next pages — site / 030122211101 / 6

- [advertise_custom.advertise_where.site.site](data-sources--udp_loadbalancer--reference--group-001.md#canonical-0132212221220101-0133210003132200-0222131122221002-0112021001222330-2230231331213111-1230020000100322-1103313030020131-3013030130203221)
- [advertise_custom.advertise_where](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2330333120302301-0122113032301122-1110222000313131-0313010302030110-0102002133220333-1223323113301222-1013010221110031-2001020101120113)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)

<a id="canonical-0132212221220101-0133210003132200-0222131122221002-0112021001222330-2230231331213111-1230020000100322-1103313030020131-3013030130203221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3233203000202232-0200102201201031-2120213110033332-1333133320112331-2212313133202212-0313120330223131-1030023003121213-1202130100320012"></a>

## advertise_custom.advertise_where.site.site — site / 303301123011 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2100233330113013-1011020101321013-1113020011303120-3332020200220223-1130023001010003-0211031020201213-0121011023231220-3020101330032130)
- [advertise_custom](data-sources--udp_loadbalancer--reference--group-001.md#canonical-0102011322232323-2032232030221030-3013012213003033-1131020131301123-2120010001223303-0101100312231222-3210322122303222-2321103220212112)
- [advertise_custom.advertise_where](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2330333120302301-0122113032301122-1110222000313131-0313010302030110-0102002133220333-1223323113301222-1013010221110031-2001020101120113)
- [advertise_custom.advertise_where.site](data-sources--udp_loadbalancer--reference--group-001.md#canonical-0001321010020213-0203311323323110-0233120231122002-3331103102133122-0331330110330133-3010302003030201-3012012310300300-0310212103130331)
- advertise_custom.advertise_where.site.site

<a id="canonical-1133221030200030-0000311013101303-3022330032230202-1203032023021322-1333303333023322-2121012233032301-2121231111013310-2101030201231120"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-1211012131001211-3112221020032021-1100020102001200-3331203201011033-3220112133231013-1201122111222302-2232033332222220-1031010132230221"></a>

## Direct properties — site / 303301123011 / 3

<a id="canonical-0031202101313013-3321223100331100-0132302310101013-0002211122221323-2310023211231013-1220112231232320-2232202023100323-2312120131020333"></a>

<a id="canonical-1011112222133220-1023001200122110-3132301020133313-1113322330210202-1033100011303032-1120303133022202-2112213130312000-2112232312331203"></a>

## name property — site / 303301123011 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-3113111133302232-0011011133012331-3000231022021300-1312131331010132-2213202020333313-0121103322111021-1313330032330022-2202020122032220"></a>

<a id="canonical-3032222103232122-2121322203111300-1323220010033220-1333111232213112-3033201320311231-1133023132131103-0223331011222212-0112133000232310"></a>

## namespace property — site / 303301123011 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
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
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3213123021132100-1303223233203333-2321213201231223-1303011300233322-3201031132232222-1201331322222210-2222210302032321-2130131201322011"></a>

<a id="canonical-3321021003310032-1311233323002203-1000202323011030-0233300030022302-3211313223322311-3102223100032030-3231302320010020-2113220101123001"></a>

## tenant property — site / 303301123011 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-2220001212022330-2332102033030111-0113221100300100-2122010210330033-1120020123231121-3221110003130102-2012001330100023-3332123232011110"></a>

## Next pages — site / 303301123011 / 7

- [advertise_custom.advertise_where.site](data-sources--udp_loadbalancer--reference--group-001.md#canonical-0001321010020213-0203311323323110-0233120231122002-3331103102133122-0331330110330133-3010302003030201-3012012310300300-0310212103130331)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)

<a id="canonical-2100122031331002-2003002020311110-1200030312012302-1130213020220111-1103122201220211-0000212121123211-0222320133113302-0200220321012100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1212301011120330-2112221230303013-0300103001210113-2312201101332102-1113023113101131-1000332330100033-0112322331021023-3130321303202231"></a>

## advertise_custom.advertise_where.use_default_port — use_default_port / 202211332013 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2100233330113013-1011020101321013-1113020011303120-3332020200220223-1130023001010003-0211031020201213-0121011023231220-3020101330032130)
- [advertise_custom](data-sources--udp_loadbalancer--reference--group-001.md#canonical-0102011322232323-2032232030221030-3013012213003033-1131020131301123-2120010001223303-0101100312231222-3210322122303222-2321103220212112)
- [advertise_custom.advertise_where](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2330333120302301-0122113032301122-1110222000313131-0313010302030110-0102002133220333-1223323113301222-1013010221110031-2001020101120113)
- advertise_custom.advertise_where.use_default_port

<a id="canonical-2321330311032101-3033001320102311-0030323100101321-2012203213002110-2011102032333220-1002332023300233-1320220223333021-0303301113210221"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

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

<a id="canonical-0212100203211131-0312112102101230-2232001203112010-0311222131223102-0011033121033003-0233012222312132-3232331132330013-1002212031230120"></a>

## Direct properties — use_default_port / 202211332013 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0131121321010301-3212021100320112-1232123200222121-2110200030331223-0032021110230302-1320012122332011-1233300001312113-0221210200212123"></a>

## Next pages — use_default_port / 202211332013 / 4

- [advertise_custom.advertise_where](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2330333120302301-0122113032301122-1110222000313131-0313010302030110-0102002133220333-1223323113301222-1013010221110031-2001020101120113)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)

<a id="canonical-3220100322102130-0323221211203111-0221213211203323-1020220322130200-0230320321133323-3220103132032220-0112232113302131-3013003122012013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2320211002210111-2023232010001023-1122210013212111-3330331302031232-3133100121312112-1020233022202311-3030231111323120-0000203230033220"></a>

## advertise_custom.advertise_where.virtual_network — virtual_network / 301203233230 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2100233330113013-1011020101321013-1113020011303120-3332020200220223-1130023001010003-0211031020201213-0121011023231220-3020101330032130)
- [advertise_custom](data-sources--udp_loadbalancer--reference--group-001.md#canonical-0102011322232323-2032232030221030-3013012213003033-1131020131301123-2120010001223303-0101100312231222-3210322122303222-2321103220212112)
- [advertise_custom.advertise_where](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2330333120302301-0122113032301122-1110222000313131-0313010302030110-0102002133220333-1223323113301222-1013010221110031-2001020101120113)
- advertise_custom.advertise_where.virtual_network

<a id="canonical-0301200312031021-3021021033202132-1213111301123303-1031113310023222-3103322332311103-0330312223003203-0333111030310133-3031132203131023"></a>

Type: `"single"`. Computed.

Parameters to advertise on a given virtual network.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-v6_vip_choice": "[\"default_v6_vip\",\"specific_v6_vip\"]",
  "x-ves-oneof-field-vip_choice": "[\"default_vip\",\"specific_vip\"]"
}
```

<a id="canonical-3310300011300232-1220313020013003-3322121333031320-1111203112123222-3211012130303303-2200211020010031-2020103033301313-2211330101200021"></a>

## Direct properties — virtual_network / 301203233230 / 3

- [default_v6_vip](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1023130000112310-2213211000322303-0311032330330113-1130130223010023-1022033032001033-1013203120223030-1010001022313033-2330102033202102): complete subsection reference.

- [default_vip](data-sources--udp_loadbalancer--reference--group-001.md#canonical-3201032330300103-1301121331222120-3112202323120220-1233113130010001-3133112032101310-1200012020230131-1212331101031210-2310220232012131): complete subsection reference.

<a id="canonical-0103003003133020-3010331332222230-2330030130013220-0110300013313110-2131113230002103-2023033302222201-3100022300313121-0011210202201000"></a>

<a id="canonical-3221210203102021-0101110112103112-1110232330030312-3120200110231321-3233220030113230-2322322031112121-2202220110121130-1131000322100020"></a>

## specific_v6_vip property — virtual_network / 301203233230 / 4

Type: `"string"`. Computed.

Exclusive with \[default\_v6\_vip\] Use given IPv6 address as VIP on virtual Network.

Upstream description:

Exclusive with \[default\_v6\_vip\] Use given IPv6 address as VIP on virtual Network.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-3212202212220323-3301202332222032-0022311333031220-3301203133133120-0010310201020033-1233331031002122-1000100000312130-1223200302300113"></a>

<a id="canonical-3231100320302023-0130102011312030-2233102302320300-3002110030222300-2213313030323332-2132302130302211-2000013331223120-2132031002012322"></a>

## specific_vip property — virtual_network / 301203233230 / 5

Type: `"string"`. Computed.

Exclusive with \[default\_vip\] Use given IPv4 address as VIP on virtual Network.

Upstream description:

Exclusive with \[default\_vip\] Use given IPv4 address as VIP on virtual Network.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

- [virtual_network](data-sources--udp_loadbalancer--reference--group-001.md#canonical-3020302033022121-3233222030313121-1313031203322111-1221101212033101-1321201111021303-3011210122023320-1200313332212021-2323201322012213): complete subsection reference.

<a id="canonical-0320201330212332-3003003130203102-1300021011003213-3301300130333221-3012330303211332-0131133012032211-3021331002121221-3021332021032333"></a>

## Next pages — virtual_network / 301203233230 / 6

- [advertise_custom.advertise_where.virtual_network.default_v6_vip](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1023130000112310-2213211000322303-0311032330330113-1130130223010023-1022033032001033-1013203120223030-1010001022313033-2330102033202102)
- [advertise_custom.advertise_where.virtual_network.default_vip](data-sources--udp_loadbalancer--reference--group-001.md#canonical-3201032330300103-1301121331222120-3112202323120220-1233113130010001-3133112032101310-1200012020230131-1212331101031210-2310220232012131)
- [advertise_custom.advertise_where.virtual_network.virtual_network](data-sources--udp_loadbalancer--reference--group-001.md#canonical-3020302033022121-3233222030313121-1313031203322111-1221101212033101-1321201111021303-3011210122023320-1200313332212021-2323201322012213)
- [advertise_custom.advertise_where](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2330333120302301-0122113032301122-1110222000313131-0313010302030110-0102002133220333-1223323113301222-1013010221110031-2001020101120113)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)

<a id="canonical-1023130000112310-2213211000322303-0311032330330113-1130130223010023-1022033032001033-1013203120223030-1010001022313033-2330102033202102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1111322012030312-1330202332030100-3113131330221020-1133322032330332-1210130232113331-0220130121032311-0311211022123210-3111010232332301"></a>

## advertise_custom.advertise_where.virtual_network.default_v6_vip — default_v6_vip / 231020132203 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2100233330113013-1011020101321013-1113020011303120-3332020200220223-1130023001010003-0211031020201213-0121011023231220-3020101330032130)
- [advertise_custom](data-sources--udp_loadbalancer--reference--group-001.md#canonical-0102011322232323-2032232030221030-3013012213003033-1131020131301123-2120010001223303-0101100312231222-3210322122303222-2321103220212112)
- [advertise_custom.advertise_where](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2330333120302301-0122113032301122-1110222000313131-0313010302030110-0102002133220333-1223323113301222-1013010221110031-2001020101120113)
- [advertise_custom.advertise_where.virtual_network](data-sources--udp_loadbalancer--reference--group-001.md#canonical-3220100322102130-0323221211203111-0221213211203323-1020220322130200-0230320321133323-3220103132032220-0112232113302131-3013003122012013)
- advertise_custom.advertise_where.virtual_network.default_v6_vip

<a id="canonical-2003113332311120-0003321011000110-3223103020300123-2320200001221012-1222212311130230-3001332323120020-0001311113132300-1213011223333211"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

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

<a id="canonical-1223003121323220-1202111220233331-2021000123010300-3131322020033212-3322203212031213-2212131110203023-1133133201201333-2112310200213313"></a>

## Direct properties — default_v6_vip / 231020132203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2111230030301000-1010222312131112-2033231130310203-1013022231210001-3331113213233301-3023001030233130-3302322033321213-2121222111210113"></a>

## Next pages — default_v6_vip / 231020132203 / 4

- [advertise_custom.advertise_where.virtual_network](data-sources--udp_loadbalancer--reference--group-001.md#canonical-3220100322102130-0323221211203111-0221213211203323-1020220322130200-0230320321133323-3220103132032220-0112232113302131-3013003122012013)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)

<a id="canonical-3201032330300103-1301121331222120-3112202323120220-1233113130010001-3133112032101310-1200012020230131-1212331101031210-2310220232012131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2121033231322132-0112100021221103-1003012010310133-1230012330320011-2333202213312023-1311033011203020-1210223201233031-1300102120030120"></a>

## advertise_custom.advertise_where.virtual_network.default_vip — default_vip / 030300230103 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2100233330113013-1011020101321013-1113020011303120-3332020200220223-1130023001010003-0211031020201213-0121011023231220-3020101330032130)
- [advertise_custom](data-sources--udp_loadbalancer--reference--group-001.md#canonical-0102011322232323-2032232030221030-3013012213003033-1131020131301123-2120010001223303-0101100312231222-3210322122303222-2321103220212112)
- [advertise_custom.advertise_where](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2330333120302301-0122113032301122-1110222000313131-0313010302030110-0102002133220333-1223323113301222-1013010221110031-2001020101120113)
- [advertise_custom.advertise_where.virtual_network](data-sources--udp_loadbalancer--reference--group-001.md#canonical-3220100322102130-0323221211203111-0221213211203323-1020220322130200-0230320321133323-3220103132032220-0112232113302131-3013003122012013)
- advertise_custom.advertise_where.virtual_network.default_vip

<a id="canonical-0010220213213022-0013121302323110-0110210311302020-2123323312320210-0230111313011130-3030132020020232-1103311122321112-2221211312221200"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

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

<a id="canonical-2130303210320001-3011212211221112-3300011300112003-3031230031102302-1300213210023000-3102200021013230-2002031211232213-3110212010322012"></a>

## Direct properties — default_vip / 030300230103 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1220023111312030-1200021132300331-2110110222120113-0210332203130022-3021010323230023-0221033020211001-1010112001311201-2011002212332221"></a>

## Next pages — default_vip / 030300230103 / 4

- [advertise_custom.advertise_where.virtual_network](data-sources--udp_loadbalancer--reference--group-001.md#canonical-3220100322102130-0323221211203111-0221213211203323-1020220322130200-0230320321133323-3220103132032220-0112232113302131-3013003122012013)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)

<a id="canonical-3020302033022121-3233222030313121-1313031203322111-1221101212033101-1321201111021303-3011210122023320-1200313332212021-2323201322012213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1323231300232331-0310102232213333-0313033012001300-0203213031000011-3113300101232321-0331020302002230-1001330303122030-2200212203323312"></a>

## advertise_custom.advertise_where.virtual_network.virtual_network — virtual_network / 111131231311 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2100233330113013-1011020101321013-1113020011303120-3332020200220223-1130023001010003-0211031020201213-0121011023231220-3020101330032130)
- [advertise_custom](data-sources--udp_loadbalancer--reference--group-001.md#canonical-0102011322232323-2032232030221030-3013012213003033-1131020131301123-2120010001223303-0101100312231222-3210322122303222-2321103220212112)
- [advertise_custom.advertise_where](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2330333120302301-0122113032301122-1110222000313131-0313010302030110-0102002133220333-1223323113301222-1013010221110031-2001020101120113)
- [advertise_custom.advertise_where.virtual_network](data-sources--udp_loadbalancer--reference--group-001.md#canonical-3220100322102130-0323221211203111-0221213211203323-1020220322130200-0230320321133323-3220103132032220-0112232113302131-3013003122012013)
- advertise_custom.advertise_where.virtual_network.virtual_network

<a id="canonical-2100120132222112-1333120200020002-3323021130233113-0201032111113232-0023030221111211-0022132322121323-2121011221312031-3021100102101122"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-0102230012301202-2031223032033101-1131012002121003-0330002331202322-2210323331030012-1131101132131123-2220212111022013-3100000103020012"></a>

## Direct properties — virtual_network / 111131231311 / 3

<a id="canonical-2131000211310123-3023331330303130-0330303302022002-1131123122300100-3312131013232013-3303213001133020-1212033221322303-1310321212103223"></a>

<a id="canonical-3301302100011112-0023022113013320-1330200113213301-2131100012130212-2031023112011121-0313202201302023-3333332011003130-0333211012111131"></a>

## name property — virtual_network / 111131231311 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-0122122222321103-3033010301011201-1022011003333021-2121121233103232-2122030112112303-3212032022333301-3330310231303233-0122022100232322"></a>

<a id="canonical-0001300031210233-1321020330211031-2033131312121001-3022003232232230-1102012003122101-0303320323212033-3000023332301003-2302313210202323"></a>

## namespace property — virtual_network / 111131231311 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
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
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1020010113013302-1110120131020231-1301121001111132-3223302032311221-1000332012001102-3002130222130133-3203300003133113-3021120223123303"></a>

<a id="canonical-3330102211100133-0231111131112210-1022311001223313-3022203313001321-3032100223011203-2121002310220020-0111101331302111-1232111033313300"></a>

## tenant property — virtual_network / 111131231311 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-0331310121003003-0110123230013110-3320120112023302-1123110332030122-0333303301201332-2321303203101212-3213313323300321-3011301211232200"></a>

## Next pages — virtual_network / 111131231311 / 7

- [advertise_custom.advertise_where.virtual_network](data-sources--udp_loadbalancer--reference--group-001.md#canonical-3220100322102130-0323221211203111-0221213211203323-1020220322130200-0230320321133323-3220103132032220-0112232113302131-3013003122012013)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)

<a id="canonical-3203131010231322-1001331213122003-3210202200131011-3130020031032230-3031322232223100-3011212301111021-0133201102201212-1303301012202311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0123013303333200-1003232221233302-3012001122121301-3011013020300332-1003100131130231-1302001002212203-0131012013203332-0201021313110003"></a>

## advertise_custom.advertise_where.virtual_site — virtual_site / 333231110101 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2100233330113013-1011020101321013-1113020011303120-3332020200220223-1130023001010003-0211031020201213-0121011023231220-3020101330032130)
- [advertise_custom](data-sources--udp_loadbalancer--reference--group-001.md#canonical-0102011322232323-2032232030221030-3013012213003033-1131020131301123-2120010001223303-0101100312231222-3210322122303222-2321103220212112)
- [advertise_custom.advertise_where](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2330333120302301-0122113032301122-1110222000313131-0313010302030110-0102002133220333-1223323113301222-1013010221110031-2001020101120113)
- advertise_custom.advertise_where.virtual_site

<a id="canonical-1122203022113201-1112121013120332-0003203300331102-1232022311123213-3202002211022112-0023020000300013-0121220301323311-2302031110310311"></a>

Type: `"single"`. Computed.

Defines a reference to a customer site virtual site along with network type where a load balancer
could be advertised.

Upstream description:

This defines a reference to a customer site virtual site along with network type where a load
balancer could be advertised.

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

<a id="canonical-3001333100213231-0310221003010121-2002212302000331-1010133302202201-3231323221303220-1111201010033211-0121332121031131-1333030132330102"></a>

## Direct properties — virtual_site / 333231110101 / 3

<a id="canonical-1213123113003300-2011031302230033-3123030212113130-2010211233323311-1221302303332230-1021122030130203-3132203210312111-2222101132002110"></a>

<a id="canonical-2310003302301230-1030313100320103-2301201023121332-1231111023323101-3232131201103333-0323212322113030-2220203020330110-2223232203122201"></a>

## network property — virtual_site / 333231110101 / 4

Type: `"string"`. Computed.

\[Enum:
SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE|SITE\_NETWORK\_INSIDE|SITE\_NETWORK\_OUTSIDE|SITE\_NETWORK\_SERVICE|SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_IP\_FABRIC\]
Defines network types to be used on site All inside and outside networks. All inside and outside
networks with internet VIP support. All inside networks. Possible values are
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`, \`SITE\_NETWORK\_INSIDE\`, \`SITE\_NETWORK\_OUTSIDE\`,
\`SITE\_NETWORK\_SERVICE\`, \`SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP\`,
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP\`, \`SITE\_NETWORK\_IP\_FABRIC\`.
Defaults to \`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`.

Upstream description:

This defines network types to be used on site

All inside and outside networks. All inside and outside networks with internet VIP support. All
inside networks. All outside networks. All outside networks with internet VIP support. VK8s service
network. &#8203;- SITE\_NETWORK\_IP\_FABRIC: VER IP Fabric network for the site

This Virtual network type is used for exposing virtual host on IP Fabric network on the VER site or
for endpoint in IP Fabric network.

Receipt-pinned upstream constraints:

```json
{
  "default": "SITE_NETWORK_INSIDE_AND_OUTSIDE",
  "enum": [
    "SITE_NETWORK_INSIDE_AND_OUTSIDE",
    "SITE_NETWORK_INSIDE",
    "SITE_NETWORK_OUTSIDE",
    "SITE_NETWORK_SERVICE",
    "SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_IP_FABRIC"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [virtual_site](data-sources--udp_loadbalancer--reference--group-001.md#canonical-3113210122002101-0102222101231331-2333133201132323-3000231231011321-0213212333011300-2132221021303202-1013233111113123-0201130313101023): complete subsection reference.

<a id="canonical-0311322310100210-0323103302332300-2221002022202303-0002212333103101-3001031031320123-3011020030122103-1210232233302310-0113231000033232"></a>

## Next pages — virtual_site / 333231110101 / 5

- [advertise_custom.advertise_where.virtual_site.virtual_site](data-sources--udp_loadbalancer--reference--group-001.md#canonical-3113210122002101-0102222101231331-2333133201132323-3000231231011321-0213212333011300-2132221021303202-1013233111113123-0201130313101023)
- [advertise_custom.advertise_where](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2330333120302301-0122113032301122-1110222000313131-0313010302030110-0102002133220333-1223323113301222-1013010221110031-2001020101120113)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)

<a id="canonical-3113210122002101-0102222101231331-2333133201132323-3000231231011321-0213212333011300-2132221021303202-1013233111113123-0201130313101023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2330310232331011-1231003022131132-2123012110021031-2112210222203200-0321331313103023-0112132103222212-2321311010033332-1031210013201321"></a>

## advertise_custom.advertise_where.virtual_site.virtual_site — virtual_site / 121030030301 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2100233330113013-1011020101321013-1113020011303120-3332020200220223-1130023001010003-0211031020201213-0121011023231220-3020101330032130)
- [advertise_custom](data-sources--udp_loadbalancer--reference--group-001.md#canonical-0102011322232323-2032232030221030-3013012213003033-1131020131301123-2120010001223303-0101100312231222-3210322122303222-2321103220212112)
- [advertise_custom.advertise_where](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2330333120302301-0122113032301122-1110222000313131-0313010302030110-0102002133220333-1223323113301222-1013010221110031-2001020101120113)
- [advertise_custom.advertise_where.virtual_site](data-sources--udp_loadbalancer--reference--group-001.md#canonical-3203131010231322-1001331213122003-3210202200131011-3130020031032230-3031322232223100-3011212301111021-0133201102201212-1303301012202311)
- advertise_custom.advertise_where.virtual_site.virtual_site

<a id="canonical-2122203132110010-1231323300323021-0102022013333331-0132300203132330-1211133313232201-1030013232132003-2121333232200202-3313030102012203"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-3323101122010030-2221220132021013-3302123211131303-0331112301023302-0312311013013332-2203023022123323-2102112230112322-3113320321232011"></a>

## Direct properties — virtual_site / 121030030301 / 3

<a id="canonical-2212112132232010-3110203312111013-2123013202311100-0332231003322100-0122322023121103-0311320132000002-3103232120223132-3110012320313003"></a>

<a id="canonical-2222212121302312-3131122221023003-1030231302220123-2200220203221020-1023101121203120-1022200230202031-0202022000023332-3131310111333201"></a>

## name property — virtual_site / 121030030301 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-3032232022301220-1213230032330330-1111301331220032-0303203132111202-3202021310010302-1320212030332203-3011011031221220-0320110231321300"></a>

<a id="canonical-2303210221220320-1133200330131330-1003302210102133-1030320131312220-0203222033310301-0020300311321011-1232312303223031-3012123012220133"></a>

## namespace property — virtual_site / 121030030301 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
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
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-2132021313230221-2311301123320230-2232003031013032-3132310000033112-2223210100133013-0003110122230202-1122123302023221-1233100223113113"></a>

<a id="canonical-1222011212023011-0213122013030210-1323000221210311-0332001123312123-0211013330202121-2323131123110122-1111223131022321-3212202003000203"></a>

## tenant property — virtual_site / 121030030301 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3023323301100013-0320220103021330-2002002122320311-1303321303103022-2310121210002312-0030001320333010-1313003013230310-1233303032323212"></a>

## Next pages — virtual_site / 121030030301 / 7

- [advertise_custom.advertise_where.virtual_site](data-sources--udp_loadbalancer--reference--group-001.md#canonical-3203131010231322-1001331213122003-3210202200131011-3130020031032230-3031322232223100-3011212301111021-0133201102201212-1303301012202311)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)

<a id="canonical-2212122103220300-0203010212033312-1231031331121312-1122332300022302-1012303121233033-3302333231031311-1332110132333011-2300133232310310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3312230013103331-3333332223333310-0022130010000130-3130303212222001-1133011301331203-1003132011120022-2133313031320030-3223303120331123"></a>

## advertise_custom.advertise_where.virtual_site_with_vip — virtual_site_with_vip / 223100222022 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2100233330113013-1011020101321013-1113020011303120-3332020200220223-1130023001010003-0211031020201213-0121011023231220-3020101330032130)
- [advertise_custom](data-sources--udp_loadbalancer--reference--group-001.md#canonical-0102011322232323-2032232030221030-3013012213003033-1131020131301123-2120010001223303-0101100312231222-3210322122303222-2321103220212112)
- [advertise_custom.advertise_where](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2330333120302301-0122113032301122-1110222000313131-0313010302030110-0102002133220333-1223323113301222-1013010221110031-2001020101120113)
- advertise_custom.advertise_where.virtual_site_with_vip

<a id="canonical-0330110333112131-1121203013212202-2330230233030330-3111220113210323-3320201312003312-0210331330331031-1012210002313213-2232223320113002"></a>

Type: `"single"`. Computed.

Defines a reference to a customer site virtual site along with network type and IP where a load
balancer could be advertised.

Upstream description:

This defines a reference to a customer site virtual site along with network type and IP where a load
balancer could be advertised.

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

<a id="canonical-1012232220020031-0003200122010322-0200333013232322-0122202211003220-3333211223311121-3321030210133213-1133012103300001-0110212101322210"></a>

## Direct properties — virtual_site_with_vip / 223100222022 / 3

<a id="canonical-3332333313331103-0322312213002022-0121220102123300-2211321202313331-1201300233112332-2311001032312103-2322222012023100-2020332133312303"></a>

<a id="canonical-1221221122310220-1100010113120201-0020011021032021-3023333213203203-2313212010321311-0202302130011001-0020301210302133-1223222210311012"></a>

## ip property — virtual_site_with_vip / 223100222022 / 4

Type: `"string"`. Computed.

Use given IP address as VIP on the site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-1232111333201111-1031132020003112-0202223111033313-0211120220323031-0213002323223031-1130023120200321-1223310020213122-3132003000231000"></a>

<a id="canonical-1220133002003233-3310123003220231-2222300023302121-0223031303213002-1331201003132301-2122002001112013-3013331200321311-2332023110011330"></a>

## network property — virtual_site_with_vip / 223100222022 / 5

Type: `"string"`. Computed.

\[Enum: SITE\_NETWORK\_SPECIFIED\_VIP\_OUTSIDE|SITE\_NETWORK\_SPECIFIED\_VIP\_INSIDE\] Defines
network types to be used on virtual-site with specified VIP All outside networks. All inside
networks. Possible values are \`SITE\_NETWORK\_SPECIFIED\_VIP\_OUTSIDE\`,
\`SITE\_NETWORK\_SPECIFIED\_VIP\_INSIDE\`. Defaults to \`SITE\_NETWORK\_SPECIFIED\_VIP\_OUTSIDE\`.

Upstream description:

This defines network types to be used on virtual-site with specified VIP

All outside networks. All inside networks.

Receipt-pinned upstream constraints:

```json
{
  "default": "SITE_NETWORK_SPECIFIED_VIP_OUTSIDE",
  "enum": [
    "SITE_NETWORK_SPECIFIED_VIP_OUTSIDE",
    "SITE_NETWORK_SPECIFIED_VIP_INSIDE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [virtual_site](data-sources--udp_loadbalancer--reference--group-001.md#canonical-3130302313123333-2102111020031310-0223321011211222-3032113203333210-2331103130120113-0310220100103211-3123013232332321-2202010330231332): complete subsection reference.

<a id="canonical-1210123103322202-1112030231033002-1130213020103121-1021311003233022-0011330233112301-0112311302321121-0203230313201020-3001110323132122"></a>

## Next pages — virtual_site_with_vip / 223100222022 / 6

- [advertise_custom.advertise_where.virtual_site_with_vip.virtual_site](data-sources--udp_loadbalancer--reference--group-001.md#canonical-3130302313123333-2102111020031310-0223321011211222-3032113203333210-2331103130120113-0310220100103211-3123013232332321-2202010330231332)
- [advertise_custom.advertise_where](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2330333120302301-0122113032301122-1110222000313131-0313010302030110-0102002133220333-1223323113301222-1013010221110031-2001020101120113)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)

<a id="canonical-3130302313123333-2102111020031310-0223321011211222-3032113203333210-2331103130120113-0310220100103211-3123013232332321-2202010330231332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0103023013011110-1112310002003231-3020200322110312-1013023111223022-1130121233212312-2120300022333220-2131310220130031-2003303331300303"></a>

## advertise_custom.advertise_where.virtual_site_with_vip.virtual_site — virtual_site / 012000313322 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2100233330113013-1011020101321013-1113020011303120-3332020200220223-1130023001010003-0211031020201213-0121011023231220-3020101330032130)
- [advertise_custom](data-sources--udp_loadbalancer--reference--group-001.md#canonical-0102011322232323-2032232030221030-3013012213003033-1131020131301123-2120010001223303-0101100312231222-3210322122303222-2321103220212112)
- [advertise_custom.advertise_where](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2330333120302301-0122113032301122-1110222000313131-0313010302030110-0102002133220333-1223323113301222-1013010221110031-2001020101120113)
- [advertise_custom.advertise_where.virtual_site_with_vip](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2212122103220300-0203010212033312-1231031331121312-1122332300022302-1012303121233033-3302333231031311-1332110132333011-2300133232310310)
- advertise_custom.advertise_where.virtual_site_with_vip.virtual_site

<a id="canonical-1233022300012223-0303033303220000-3320010233033320-0002310033022002-3130032311102100-1203331232300120-2223132213111102-0232031110302101"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-0121223213113120-2200302211321131-3010001033320132-3012303110330333-2303030100212333-3101033031332012-1023020132210113-0020120201222103"></a>

## Direct properties — virtual_site / 012000313322 / 3

<a id="canonical-2022313133110032-1020300321310110-0023010212200233-1023320021010032-0330132003302231-1302210030123002-0211022212030202-3212112210310211"></a>

<a id="canonical-0231023201202100-0002030311111301-3212320023032121-3223322132133302-0102102202110313-3321111123210303-2001022332223133-3112331201120103"></a>

## name property — virtual_site / 012000313322 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-2230200232211301-3111120110132202-1311011231212212-1300213020203013-0200221001112113-2111001301121132-2100332112111121-0032213311012323"></a>

<a id="canonical-1203122033203021-3110211321100223-2001110323032131-3300011223330113-0212301030331020-0222101100223331-1122010100131022-0130231201131203"></a>

## namespace property — virtual_site / 012000313322 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
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
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1111120102123103-1020332320103333-2012202323121201-2111322113212023-0112220030131323-2032131003231132-1232231313202002-3301213122121311"></a>

<a id="canonical-1221303222102020-0302021130232131-3110222133211132-3010312312332110-1012000333231022-0331220203012131-0111130002013120-0322212302100010"></a>

## tenant property — virtual_site / 012000313322 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1222210200031032-2123123033313303-1200023010033231-0210301120302130-3323012102221000-3323000020123213-0011333320012002-3212111012102322"></a>

## Next pages — virtual_site / 012000313322 / 7

- [advertise_custom.advertise_where.virtual_site_with_vip](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2212122103220300-0203010212033312-1231031331121312-1122332300022302-1012303121233033-3302333231031311-1332110132333011-2300133232310310)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)

<a id="canonical-1322123110212212-1033222023310221-1021233002213200-2121220132323000-1120320111033002-2012002311123220-0020131200103231-0323132202303320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0002002213012320-2311301121030012-1032031001101301-3303233320032013-1220111133033001-1313333110303302-0111201302112223-2323232033032211"></a>

## advertise_custom.advertise_where.vk8s_service — vk8s_service / 233131212230 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2100233330113013-1011020101321013-1113020011303120-3332020200220223-1130023001010003-0211031020201213-0121011023231220-3020101330032130)
- [advertise_custom](data-sources--udp_loadbalancer--reference--group-001.md#canonical-0102011322232323-2032232030221030-3013012213003033-1131020131301123-2120010001223303-0101100312231222-3210322122303222-2321103220212112)
- [advertise_custom.advertise_where](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2330333120302301-0122113032301122-1110222000313131-0313010302030110-0102002133220333-1223323113301222-1013010221110031-2001020101120113)
- advertise_custom.advertise_where.vk8s_service

<a id="canonical-2102013110032202-3001232200131133-3300213031202022-3003323113120331-3130132231012201-2230100031023001-2130113211011313-3311230111330001"></a>

Type: `"single"`. Computed.

Defines a reference to a RE site or virtual site where a load balancer could be advertised in the
vK8s service network.

Upstream description:

This defines a reference to a RE site or virtual site where a load balancer could be advertised in
the vK8s service network.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"site\",\"virtual_site\"]"
}
```

<a id="canonical-1123120300320032-2230131130021230-3323101302223121-0221211232012212-3201323101302132-1002130331003132-1213033000212200-3303333011011100"></a>

## Direct properties — vk8s_service / 233131212230 / 3

- [site](data-sources--udp_loadbalancer--reference--group-001.md#canonical-3123200022320002-0030312103233102-2112112212001321-1320000013020003-2122112111232212-0210030120313100-0110133222000010-3302111132223331): complete subsection reference.

- [virtual_site](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2031023201330000-1132230103302131-2012333211200012-1203321302332330-3122033230310222-0013213203231300-3212012112211331-3013002332100323): complete subsection reference.

<a id="canonical-0133110233031021-3132003110120310-3323113111331011-1130100032100113-0113333111213101-2113012212103122-0212232003010301-0232021320102011"></a>

## Next pages — vk8s_service / 233131212230 / 4

- [advertise_custom.advertise_where.vk8s_service.site](data-sources--udp_loadbalancer--reference--group-001.md#canonical-3123200022320002-0030312103233102-2112112212001321-1320000013020003-2122112111232212-0210030120313100-0110133222000010-3302111132223331)
- [advertise_custom.advertise_where.vk8s_service.virtual_site](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2031023201330000-1132230103302131-2012333211200012-1203321302332330-3122033230310222-0013213203231300-3212012112211331-3013002332100323)
- [advertise_custom.advertise_where](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2330333120302301-0122113032301122-1110222000313131-0313010302030110-0102002133220333-1223323113301222-1013010221110031-2001020101120113)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)

<a id="canonical-3123200022320002-0030312103233102-2112112212001321-1320000013020003-2122112111232212-0210030120313100-0110133222000010-3302111132223331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1330020320033010-1023212322203133-2333201323200232-3310021133020223-0121213300321330-0300311111201203-3130032333300102-2221322231213130"></a>

## advertise_custom.advertise_where.vk8s_service.site — site / 203330122201 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2100233330113013-1011020101321013-1113020011303120-3332020200220223-1130023001010003-0211031020201213-0121011023231220-3020101330032130)
- [advertise_custom](data-sources--udp_loadbalancer--reference--group-001.md#canonical-0102011322232323-2032232030221030-3013012213003033-1131020131301123-2120010001223303-0101100312231222-3210322122303222-2321103220212112)
- [advertise_custom.advertise_where](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2330333120302301-0122113032301122-1110222000313131-0313010302030110-0102002133220333-1223323113301222-1013010221110031-2001020101120113)
- [advertise_custom.advertise_where.vk8s_service](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1322123110212212-1033222023310221-1021233002213200-2121220132323000-1120320111033002-2012002311123220-0020131200103231-0323132202303320)
- advertise_custom.advertise_where.vk8s_service.site

<a id="canonical-1131023232332023-1103321302120300-2012211012212221-1211230312030203-3313220231211331-3302220101221010-1023002101312012-2323220020031020"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-3320023232232022-1110321102003000-3112001101111322-2313232302031031-3011202212312111-0001211132300133-2011203132130103-1011222000020132"></a>

## Direct properties — site / 203330122201 / 3

<a id="canonical-2303003122203213-2010303320113203-1233323023130222-0031212113313313-1323132333003131-0332113030020131-2221232211113301-1210212230230121"></a>

<a id="canonical-2111112302113120-2300003213132023-0201211113100120-2110030031102322-1012230032223322-3310331032110010-3120122302312210-3322121113203033"></a>

## name property — site / 203330122201 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-1300102000121233-1311123331032301-0020211330113022-0022230323133032-1001211111303102-3133120132311023-3021203313000022-1001230303123311"></a>

<a id="canonical-3302132102010132-0203031211303103-1000220103012003-3233013232022122-1330120300011102-1003130022011023-2330331313120300-2001102001220233"></a>

## namespace property — site / 203330122201 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
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
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-2010223221021223-0200000210113102-0131230232311031-3310131312101301-1231113202031031-1220110310322113-1031203122233031-1232310030111132"></a>

<a id="canonical-2102301231312232-0202301010112313-2323103033321200-1122202303010100-3101231110103030-3123120221233311-2210120130123300-1222321223032011"></a>

## tenant property — site / 203330122201 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-0212300330022310-0121120133322123-0232020203222012-0211231111032120-2031202030130111-1021230203002120-0113031210302303-2213212302312001"></a>

## Next pages — site / 203330122201 / 7

- [advertise_custom.advertise_where.vk8s_service](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1322123110212212-1033222023310221-1021233002213200-2121220132323000-1120320111033002-2012002311123220-0020131200103231-0323132202303320)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)

<a id="canonical-2031023201330000-1132230103302131-2012333211200012-1203321302332330-3122033230310222-0013213203231300-3212012112211331-3013002332100323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2222332232000130-1030233013100333-3120133311232100-0102131302323331-3021302223233102-2330232322030221-0310100122011030-1130000231100013"></a>

## advertise_custom.advertise_where.vk8s_service.virtual_site — virtual_site / 211313231112 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2100233330113013-1011020101321013-1113020011303120-3332020200220223-1130023001010003-0211031020201213-0121011023231220-3020101330032130)
- [advertise_custom](data-sources--udp_loadbalancer--reference--group-001.md#canonical-0102011322232323-2032232030221030-3013012213003033-1131020131301123-2120010001223303-0101100312231222-3210322122303222-2321103220212112)
- [advertise_custom.advertise_where](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2330333120302301-0122113032301122-1110222000313131-0313010302030110-0102002133220333-1223323113301222-1013010221110031-2001020101120113)
- [advertise_custom.advertise_where.vk8s_service](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1322123110212212-1033222023310221-1021233002213200-2121220132323000-1120320111033002-2012002311123220-0020131200103231-0323132202303320)
- advertise_custom.advertise_where.vk8s_service.virtual_site

<a id="canonical-1331233233122113-2010013322120313-0103030312013322-2031330322031111-3231022203002300-1023003020222213-1303300031131111-2013300100220130"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-1120221033011031-1220231321302303-3000330313022112-0223233220231001-3001331131000101-0111302131300321-0012323330213133-2330120200133213"></a>

## Direct properties — virtual_site / 211313231112 / 3

<a id="canonical-0013102222333331-0230101033213033-1022130221212330-2002113333321102-0110301310331110-3321333130220233-2003111113012103-2033323212231202"></a>

<a id="canonical-1311131101210133-3102031323203200-2311333202102203-1222012232222030-2200303001330213-1212320213020003-0321220101331201-2002233003322132"></a>

## name property — virtual_site / 211313231112 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-1000333020313312-0210101210101211-0323333102103233-1013201122333331-1222311311202313-0233311212130013-3323003312103221-0222333122012101"></a>

<a id="canonical-0301211211301002-1121233233003003-1130322202021030-1101223030212222-0220312303223113-3121112003230132-0210121232222301-0121231012130010"></a>

## namespace property — virtual_site / 211313231112 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
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
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3221123231121320-1321201112110111-1002310333302002-3133210100001233-1333311000010232-1203010022002313-1123021033323213-2100202311103313"></a>

<a id="canonical-2010313113330212-2220110000003223-3323131131320331-1312103221300322-0223011312010331-3112303333320312-1030202023301211-1303313220130110"></a>

## tenant property — virtual_site / 211313231112 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1120222303033113-3131001212230013-0102133313210232-0012102130123123-0111002113021123-1001012200010103-3210232321120020-2211023020003301"></a>

## Next pages — virtual_site / 211313231112 / 7

- [advertise_custom.advertise_where.vk8s_service](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1322123110212212-1033222023310221-1021233002213200-2121220132323000-1120320111033002-2012002311123220-0020131200103231-0323132202303320)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)

<a id="canonical-3111011220331032-2313033231112301-0320230202101033-3101213131331212-0221323210201200-3230133203131100-3011303013112000-3303002032321131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0200020330231033-2211021120010213-3112000310302312-0230000321200013-2303022202120300-0322113211012113-1330231331122111-3222212302221032"></a>

## advertise_on_public — advertise_on_public / 322011300021 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2100233330113013-1011020101321013-1113020011303120-3332020200220223-1130023001010003-0211031020201213-0121011023231220-3020101330032130)
- advertise_on_public

<a id="canonical-3002102112331221-0230030210203322-3322133332201112-0022012021112003-1203203001313113-2100323231233212-2302013110132131-1303033330030333"></a>

Type: `"single"`. Computed.

Defines a way to advertise a load balancer on public. If optional public\_ip is provided, it will
only be advertised on RE sites where that public\_ip is available.

Upstream description:

This defines a way to advertise a load balancer on public. If optional public\_ip is provided, it
will only be advertised on RE sites where that public\_ip is available.

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

<a id="canonical-0310100103202023-1300111020012323-0101301001113101-3022310333010313-3322331122120312-3230032120300323-1100001322131010-3122102200210100"></a>

## Direct properties — advertise_on_public / 322011300021 / 3

- [public_ip](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1311133023123120-3112212032130330-2020003313320113-3112033221233123-1332321002120333-2011120222012102-2122313311113031-0022102212103010): complete subsection reference.

<a id="canonical-0303003132110221-2201223132131003-3113031231113312-0232230032101303-1031133310312311-2000021113211321-0011310311121023-3033131332231222"></a>

## Next pages — advertise_on_public / 322011300021 / 4

- [advertise_on_public.public_ip](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1311133023123120-3112212032130330-2020003313320113-3112033221233123-1332321002120333-2011120222012102-2122313311113031-0022102212103010)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2100233330113013-1011020101321013-1113020011303120-3332020200220223-1130023001010003-0211031020201213-0121011023231220-3020101330032130)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)

<a id="canonical-1311133023123120-3112212032130330-2020003313320113-3112033221233123-1332321002120333-2011120222012102-2122313311113031-0022102212103010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1122003000123213-2011123122030103-1030130133232320-1330301031001013-2303333201223003-0303000113302100-3110012212100033-0221030100020233"></a>

## advertise_on_public.public_ip — public_ip / 120330021003 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2100233330113013-1011020101321013-1113020011303120-3332020200220223-1130023001010003-0211031020201213-0121011023231220-3020101330032130)
- [advertise_on_public](data-sources--udp_loadbalancer--reference--group-001.md#canonical-3111011220331032-2313033231112301-0320230202101033-3101213131331212-0221323210201200-3230133203131100-3011303013112000-3303002032321131)
- advertise_on_public.public_ip

<a id="canonical-2023311310110331-1312203213301112-3102120203010222-2220111313321003-3202213230233321-0121232002122300-2111312131330110-3110130030033112"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-0131122113301023-3013101203201331-3303312001223200-1200310332201200-3113320112102133-0013221310013220-1111031321011000-2112311301223023"></a>

## Direct properties — public_ip / 120330021003 / 3

<a id="canonical-3002000012212032-0020030321033103-0223103033031101-1322131032203223-2310333310220303-1211003003310230-1333212100230122-0223100213332121"></a>

<a id="canonical-2303332320233220-2121101213003011-0102330022333132-3032012233302022-0210120221131103-3230120302102223-1231112022023031-0203113033020120"></a>

## name property — public_ip / 120330021003 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-0002312023330102-3132100023211031-3320003130112112-3323012020320011-3013130021221301-0001121231011231-0321021121030312-3320201211102303"></a>

<a id="canonical-2212002122211222-3033201031221233-3110012023010211-1012031300033232-0003332031030001-2122030011212300-3320011022023211-1312012012330223"></a>

## namespace property — public_ip / 120330021003 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
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
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-0130312323000223-1010300012133130-1211333203303222-2032133023130001-2010332333022122-2320312101210212-1300103332210102-1233010332202011"></a>

<a id="canonical-2203212300201313-0122133000331011-3230122223211231-1302210032023310-3103022230331002-2120032322100330-2303103322030312-2132133312221311"></a>

## tenant property — public_ip / 120330021003 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3031322322321012-2112202202003113-0233321012203332-1310120120103313-3210033120221100-1020011233011230-1210333210103122-2103022322320002"></a>

## Next pages — public_ip / 120330021003 / 7

- [advertise_on_public](data-sources--udp_loadbalancer--reference--group-001.md#canonical-3111011220331032-2313033231112301-0320230202101033-3101213131331212-0221323210201200-3230133203131100-3011303013112000-3303002032321131)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)

<a id="canonical-2311322310303223-0231020022011110-0310223232001233-0311013120223030-0200302330003100-2330123303103011-3100022130033000-2112131213103102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0020202100322031-0112102222030222-2232121213130020-3120331003023103-3320330113332001-1210222111120203-1323010122223203-3123221201302001"></a>

## advertise_on_public_default_vip — advertise_on_public_default_vip / 211311332100 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2100233330113013-1011020101321013-1113020011303120-3332020200220223-1130023001010003-0211031020201213-0121011023231220-3020101330032130)
- advertise_on_public_default_vip

<a id="canonical-2310321122112123-0123232131321211-3313201213130111-1230200200322030-3023023100331123-3301223220012322-2310311120302030-3020231201110300"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

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

<a id="canonical-1331322130331301-0110200131220032-0023011221223322-0012001222221303-2333033002001100-0322113113111001-1312010230223232-3310212031001310"></a>

## Direct properties — advertise_on_public_default_vip / 211311332100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1223121033010120-0011232210302303-0102323023310310-0211221311330312-2321201200022010-1113331212220213-0301332133131212-1210102023130230"></a>

## Next pages — advertise_on_public_default_vip / 211311332100 / 4

- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2100233330113013-1011020101321013-1113020011303120-3332020200220223-1130023001010003-0211031020201213-0121011023231220-3020101330032130)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)

<a id="canonical-0131000331300322-0103301113212111-2211110212313310-3002002230210313-3103111021003321-1003003123210201-3122100201032330-0222313230310031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1222011010301310-0120101002210000-0012010301001031-1320121023032230-2232301201032010-1310021303132101-0302023201200322-0232231233212003"></a>

## do_not_advertise — do_not_advertise / 213303311113 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2100233330113013-1011020101321013-1113020011303120-3332020200220223-1130023001010003-0211031020201213-0121011023231220-3020101330032130)
- do_not_advertise

<a id="canonical-0101121300202333-1021331103221322-2001200113112121-3023123100313000-2133032303112201-3112321013203203-1300133202010110-3233112123211320"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for do not advertise.

Upstream description:

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

<a id="canonical-2320031221120120-0310310323333100-0123103100220211-2000110031100323-0121120122003201-2031122101220220-1222231231002232-0312211030223021"></a>

## Direct properties — do_not_advertise / 213303311113 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3320331332231203-2110120303003113-3202311302320121-1122021330331222-1211223312023132-0213231313202202-1011232231210210-0322101012301330"></a>

## Next pages — do_not_advertise / 213303311113 / 4

- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2100233330113013-1011020101321013-1113020011303120-3332020200220223-1130023001010003-0211031020201213-0121011023231220-3020101330032130)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)

<a id="canonical-3002201000313203-2123320202312030-0022330303110220-0230302222121120-3201231231321301-0232331300203223-0022233321021130-0002331230323021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0302030332213130-2311033001011210-1310030030300120-3323202123031213-2320320110031301-2023303311320312-3112100200010202-1002000222131330"></a>

## hash_policy_choice_random — hash_policy_choice_random / 211132320103 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2100233330113013-1011020101321013-1113020011303120-3332020200220223-1130023001010003-0211031020201213-0121011023231220-3020101330032130)
- hash_policy_choice_random

<a id="canonical-1122223103333023-1233133133031020-1222333013212203-1211020022330121-0311223010102111-0301000123021322-3200002323120202-1133011121110231"></a>

Type: `["object", {}]`. Computed.

\[OneOf: hash\_policy\_choice\_random, hash\_policy\_choice\_round\_robin,
hash\_policy\_choice\_source\_ip\_stickiness\] Configuration parameter for hash policy choice
random.

Upstream description:

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

OneOf alternatives in this subsection:

- [hash_policy_choice_random](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1122223103333023-1233133133031020-1222333013212203-1211020022330121-0311223010102111-0301000123021322-3200002323120202-1133011121110231)
- [hash_policy_choice_round_robin](data-sources--udp_loadbalancer--reference--group-001.md#canonical-3301230302200220-2012010210111020-1011232100133300-2020330312230202-3133231231333212-0110212313312013-3023101332333311-3030201230022130)
- [hash_policy_choice_source_ip_stickiness](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2222212012222101-3221310030001111-3333212101313002-3210123130010303-1113113103110210-0132231120331121-2220203130230002-0122030111023130)

Select alternatives according to the provider validators above.

<a id="canonical-0012221122031022-1011123200223132-3133230102211303-3321013201123303-0131110213321322-3111112200320231-1012320231311223-1033303101302212"></a>

## Direct properties — hash_policy_choice_random / 211132320103 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2021002031123303-0012220101320023-2002223322133212-3003002211203332-3203302232103130-1201032323301023-1230310212101030-2021120013133222"></a>

## Next pages — hash_policy_choice_random / 211132320103 / 4

- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2100233330113013-1011020101321013-1113020011303120-3332020200220223-1130023001010003-0211031020201213-0121011023231220-3020101330032130)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)

<a id="canonical-1223232202021103-2020233233211000-0201322313330131-2301121001023022-2021122103113111-2320302310333212-3013001031012230-2333021133101031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3110012023030121-1130303312010112-3111010300133303-2131320222033301-1203013020032130-1300011003203301-1313302231321122-1010001131111102"></a>

## hash_policy_choice_round_robin — hash_policy_choice_round_robin / 112331103112 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2100233330113013-1011020101321013-1113020011303120-3332020200220223-1130023001010003-0211031020201213-0121011023231220-3020101330032130)
- hash_policy_choice_round_robin

<a id="canonical-3301230302200220-2012010210111020-1011232100133300-2020330312230202-3133231231333212-0110212313312013-3023101332333311-3030201230022130"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for hash policy choice round robin.

Upstream description:

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

<a id="canonical-1130111013003000-2130131321302023-0011222131111010-1131112031233101-2121221310011221-0300101202331223-0321102231222122-1122332322233302"></a>

## Direct properties — hash_policy_choice_round_robin / 112331103112 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1221122203313112-0031333120000112-0322333213202133-1223032300131003-0330023112023202-0333132301332110-2013331112201200-0011223202331321"></a>

## Next pages — hash_policy_choice_round_robin / 112331103112 / 4

- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2100233330113013-1011020101321013-1113020011303120-3332020200220223-1130023001010003-0211031020201213-0121011023231220-3020101330032130)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)

<a id="canonical-0312203212031200-0300032003330222-3023010030101230-1213021131301013-0011220230203121-3021031233331130-0320012111103302-1133302313101320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1221021313031313-1121013212103311-1233200230222010-2222312221211012-0223201223312231-1121131323023132-1213200023222021-1101120110311100"></a>

## hash_policy_choice_source_ip_stickiness — hash_policy_choice_source_ip_stickiness / 103113321003 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2100233330113013-1011020101321013-1113020011303120-3332020200220223-1130023001010003-0211031020201213-0121011023231220-3020101330032130)
- hash_policy_choice_source_ip_stickiness

<a id="canonical-2222212012222101-3221310030001111-3333212101313002-3210123130010303-1113113103110210-0132231120331121-2220203130230002-0122030111023130"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

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

<a id="canonical-1332312033202232-1000323010201022-2123032301331232-0003031302201301-1303020102110001-3103111232310131-3332021333313013-2130122103132012"></a>

## Direct properties — hash_policy_choice_source_ip_stickiness / 103113321003 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0113000103032301-0232202310032122-2212010102113110-0020311321100120-3011111321213022-1021202101233213-2200202020212222-2320311231131210"></a>

## Next pages — hash_policy_choice_source_ip_stickiness / 103113321003 / 4

- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2100233330113013-1011020101321013-1113020011303120-3332020200220223-1130023001010003-0211031020201213-0121011023231220-3020101330032130)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)

<a id="canonical-0311102003233330-0121133300020322-3132013220222322-3203312022330021-3033322310112312-2323110002200030-0120301333121123-3011112033313013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1003302121011300-2113132331122001-2103121022322002-0212112312023312-0233003113002030-3102020110323203-2232202031220332-2200231213031323"></a>

## no_service_policies — no_service_policies / 021002233212 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2100233330113013-1011020101321013-1113020011303120-3332020200220223-1130023001010003-0211031020201213-0121011023231220-3020101330032130)
- no_service_policies

<a id="canonical-1032213120323321-0333213321123321-0302320330333313-0032020301112220-3013111302213210-3231332120102320-2113100110303220-1220112311233030"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no service policies.

Upstream description:

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

<a id="canonical-2223220013123111-1001101300230232-2312111331233112-1120232003201313-2203112230300022-3013230300311003-2001121021313101-0211033213031300"></a>

## Direct properties — no_service_policies / 021002233212 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3032322221333130-2132023100021202-3320030302322011-0022300310103231-2311130201302110-1132010212003010-2203121203000313-3113131222101112"></a>

## Next pages — no_service_policies / 021002233212 / 4

- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2100233330113013-1011020101321013-1113020011303120-3332020200220223-1130023001010003-0211031020201213-0121011023231220-3020101330032130)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)

<a id="canonical-2320311033331002-1133320121212213-0320333113110131-2120100213202221-2233020011030301-1322013030212121-1131120221033210-2032330100122213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3021000201320312-3213021312320033-2231323010303023-0323022131021103-2230002111331312-3213211011223232-1221230130201221-3233310113330230"></a>

## origin_pools_weights — origin_pools_weights / 100100012101 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2100233330113013-1011020101321013-1113020011303120-3332020200220223-1130023001010003-0211031020201213-0121011023231220-3020101330032130)
- origin_pools_weights

<a id="canonical-0133011002012002-1223131320102012-0223032003033302-3120000021210220-3211030220302010-1303101301123100-2003300312220212-2222230333102230"></a>

Type: `"list"`. Computed.

Origin pools with weights and priorities used for this load balancer.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2303200132023030-0220312010213321-2200110332021303-3233031201022120-2212220001313212-1101102222111223-0312221123112011-2132200310332003"></a>

## Direct properties — origin_pools_weights / 100100012101 / 3

- [cluster](data-sources--udp_loadbalancer--reference--group-001.md#canonical-0230130133202333-1001133210001102-0032323221213220-3011001220031132-1111232320322003-0313021110312333-0213231231112220-1210211232312102): complete subsection reference.

- [endpoint_subsets](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1302301333211301-1123221101033131-2132303022122031-1332123230321323-3012301202031110-2010102300221033-0103213330303302-2031100012212202): complete subsection reference.

- [pool](data-sources--udp_loadbalancer--reference--group-001.md#canonical-3332332320100313-2121230201121000-2022013112331311-3201100310101211-2321012133031100-2230320330223230-3100221010222132-0132310231310320): complete subsection reference.

<a id="canonical-0300133201021131-0102133022313220-1232323133310202-3210130133113131-1110032212101112-1122320120332321-1313023210331221-3200003102111002"></a>

<a id="canonical-0230002321313302-1330013002323102-0300223223310123-2311303123111000-3220022123322310-0022100010123332-2231331232300001-2321321131020003"></a>

## priority property — origin_pools_weights / 100100012101 / 4

Type: `"number"`. Computed.

Priority of this origin pool, valid only with multiple origin pools. Value of 0 will make the pool
as lowest priority origin pool Priority of 1 means highest priority and is considered active. When
active origin pool is not available, lower priority origin pools are made active as per the..

Upstream description:

Priority of this origin pool, valid only with multiple origin pools. Value of 0 will make the pool
as lowest priority origin pool Priority of 1 means highest priority and is considered active. When
active origin pool is not available, lower priority origin pools are made active as per the
increasing priority.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-3021320002112111-3000220113303332-1010323023223120-2113122100112001-0021323333213223-2112110100301011-0111002113210303-2122311220331223"></a>

<a id="canonical-1321300001022021-0030320002303032-2223203232332020-1222332131002213-1011320033022203-3210123322202132-0233212033122322-0102133012332302"></a>

## weight property — origin_pools_weights / 100100012101 / 5

Type: `"number"`. Computed.

Weight of this origin pool, valid only with multiple origin pool. Value of 0 will disable the pool.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "load-balancing",
    "constraintType": "number",
    "maximum": 100,
    "metadata": {
      "confidence": 0.8,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0103200111122323-3312020313033101-3122011131103012-0221321211100232-0123303000222313-1000111213100101-1121300000023120-1132110022310222"></a>

## Next pages — origin_pools_weights / 100100012101 / 6

- [origin_pools_weights.cluster](data-sources--udp_loadbalancer--reference--group-001.md#canonical-0230130133202333-1001133210001102-0032323221213220-3011001220031132-1111232320322003-0313021110312333-0213231231112220-1210211232312102)
- [origin_pools_weights.endpoint_subsets](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1302301333211301-1123221101033131-2132303022122031-1332123230321323-3012301202031110-2010102300221033-0103213330303302-2031100012212202)
- [origin_pools_weights.pool](data-sources--udp_loadbalancer--reference--group-001.md#canonical-3332332320100313-2121230201121000-2022013112331311-3201100310101211-2321012133031100-2230320330223230-3100221010222132-0132310231310320)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2100233330113013-1011020101321013-1113020011303120-3332020200220223-1130023001010003-0211031020201213-0121011023231220-3020101330032130)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)

<a id="canonical-0230130133202333-1001133210001102-0032323221213220-3011001220031132-1111232320322003-0313021110312333-0213231231112220-1210211232312102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1210300201233132-0223300122023002-0103211001221232-0331121313231022-0020121022233030-2213320120233233-1311211213322313-0300012323220321"></a>

## origin_pools_weights.cluster — cluster / 131023330231 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2100233330113013-1011020101321013-1113020011303120-3332020200220223-1130023001010003-0211031020201213-0121011023231220-3020101330032130)
- [origin_pools_weights](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2320311033331002-1133320121212213-0320333113110131-2120100213202221-2233020011030301-1322013030212121-1131120221033210-2032330100122213)
- origin_pools_weights.cluster

<a id="canonical-1031211211320302-0031131212201002-0213103313332201-0310321132230101-2230211130210321-3311132102223222-2211022320312032-0203230021120210"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-0302213001333321-0223112321021222-3221310310111022-1100300333321221-3200212223201003-3211303203000222-0323032300200031-1013020320020233"></a>

## Direct properties — cluster / 131023330231 / 3

<a id="canonical-3231203300033101-0101013201221012-3020103010033332-1200322312113010-3010313223001331-1033311313102112-0312120013210303-3211113133221023"></a>

<a id="canonical-3200201311003230-3102032103310000-0312323012323100-2022222032303223-2123331023123300-1120030220303133-2030033022122033-2300321233002222"></a>

## name property — cluster / 131023330231 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-3230222231000021-2130103220030303-3223213321211322-2001132112020121-0121222113103200-1311313113131100-1102123322321312-3031221311222211"></a>

<a id="canonical-1120132100203331-0123202133313323-1011212113301103-2231033333222123-3221300221211113-2201201131122220-1003031222322131-2131003123202011"></a>

## namespace property — cluster / 131023330231 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
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
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3030322101121221-2110220133302203-0222212302322330-1000121101200122-2200303231203313-2233233102302022-0112123022030133-1120132332102320"></a>

<a id="canonical-3010032221120223-1312121230321112-2013222022032003-2202001003333020-3112323020003013-1202220110222120-0122202000303012-0233011122031113"></a>

## tenant property — cluster / 131023330231 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1220103100322130-1312132200302033-3212232120200122-3110103100202300-3000122210121321-0003310211313332-3100310200202131-3223200132131203"></a>

## Next pages — cluster / 131023330231 / 7

- [origin_pools_weights](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2320311033331002-1133320121212213-0320333113110131-2120100213202221-2233020011030301-1322013030212121-1131120221033210-2032330100122213)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)

<a id="canonical-1302301333211301-1123221101033131-2132303022122031-1332123230321323-3012301202031110-2010102300221033-0103213330303302-2031100012212202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1212132313003013-3223303322313330-3302202210001223-0320333110130212-2023113112222323-2023201313020130-3223331322311321-1131131213220020"></a>

## origin_pools_weights.endpoint_subsets — endpoint_subsets / 110100321232 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2100233330113013-1011020101321013-1113020011303120-3332020200220223-1130023001010003-0211031020201213-0121011023231220-3020101330032130)
- [origin_pools_weights](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2320311033331002-1133320121212213-0320333113110131-2120100213202221-2233020011030301-1322013030212121-1131120221033210-2032330100122213)
- origin_pools_weights.endpoint_subsets

<a id="canonical-0110132010111023-1201101220213321-0112302210211033-0112013012312031-3031003010202233-3023232011231300-0333230011331310-2213332030230103"></a>

Type: `"single"`. Computed.

Upstream origin pool may be configured to divide its origin servers into subsets based on metadata
attached to the origin servers. Routes may then specify the metadata that a endpoint must match in
order to be selected by the load balancer For origin servers which are discovered in K8s or Consul..

Upstream description:

Upstream origin pool may be configured to divide its origin servers into subsets based on metadata
attached to the origin servers. Routes may then specify the metadata that a endpoint must match in
order to be selected by the load balancer

For origin servers which are discovered in K8s or Consul cluster, the label of the service is merged
with endpoint's labels. In case of Consul, the label is derived from the "Tag" field. For labels
that are common between configured endpoint and discovered service, labels from discovered service
takes precedence.

List of key-value pairs that will be used as matching metadata. Only those origin servers of
upstream origin pool which match this metadata will be selected for load balancing.

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
    "ves.io.schema.rules.map.max_pairs": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.max_pairs": "16"
  }
}
```

<a id="canonical-2121030323311310-3102000001210232-1311231010011111-2221002210223233-1200001010313303-0120310213003221-2010033022220231-0111201102101120"></a>

## Direct properties — endpoint_subsets / 110100321232 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0210022231321132-2123332322002021-3001131320210132-0301123211223330-1111211122023112-3130312122331330-1123032103111313-1323200113312231"></a>

## Next pages — endpoint_subsets / 110100321232 / 4

- [origin_pools_weights](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2320311033331002-1133320121212213-0320333113110131-2120100213202221-2233020011030301-1322013030212121-1131120221033210-2032330100122213)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)

<a id="canonical-3332332320100313-2121230201121000-2022013112331311-3201100310101211-2321012133031100-2230320330223230-3100221010222132-0132310231310320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0121100003302101-0010130300000023-0330002310323122-1210010302310130-1101032201031312-0021200310020201-1332011311123212-0122321333133011"></a>

## origin_pools_weights.pool — pool / 013312131320 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2100233330113013-1011020101321013-1113020011303120-3332020200220223-1130023001010003-0211031020201213-0121011023231220-3020101330032130)
- [origin_pools_weights](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2320311033331002-1133320121212213-0320333113110131-2120100213202221-2233020011030301-1322013030212121-1131120221033210-2032330100122213)
- origin_pools_weights.pool

<a id="canonical-1323133101200322-1213302302201112-1032111111213103-3301221020002102-1323332123333211-1303212200321110-2330200321113122-2230233312313103"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-3211321301310211-2131232212121022-2101103121010213-1333202102222033-0122000112302321-1012012133331201-1200003202020133-2031121332332131"></a>

## Direct properties — pool / 013312131320 / 3

<a id="canonical-1230102213020221-3212202003330320-0102302113020000-1313031213130322-0311213030333120-2020110332131032-3012020223102113-0320020303212000"></a>

<a id="canonical-2000330010301111-2012331002033013-1120323233112200-2131120123231200-1122211232230220-0100000001320121-2220031201131333-1303203131212201"></a>

## name property — pool / 013312131320 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-2321200220223222-3332100322023130-1031031332012302-3330220231301033-3302103223011033-1212121010332100-1033323111013100-2320003012001213"></a>

<a id="canonical-0101222021231102-1122230210032323-0101211222110220-1210133232203232-1000131332101011-1132233032221022-3222301020233222-2233001231220311"></a>

## namespace property — pool / 013312131320 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
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
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1031331003021331-0102101320131111-0220220322021111-1212121010311210-2111000303222100-3032011112030313-2020231112122331-3113013103110130"></a>

<a id="canonical-0311223130223133-3133332010331102-0323031212221231-1311200020221101-1332111310023012-3100313322002021-2103112123133311-3320300230211201"></a>

## tenant property — pool / 013312131320 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3302103300210002-1222210003222033-1012222110333221-3222102210210121-1130201303121230-0112032031133212-1021022323321021-2311221312120110"></a>

## Next pages — pool / 013312131320 / 7

- [origin_pools_weights](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2320311033331002-1133320121212213-0320333113110131-2120100213202221-2233020011030301-1322013030212121-1131120221033210-2032330100122213)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)

<a id="canonical-3122231333000123-0312111333033323-2323101223110320-3223113220211321-0222023230110012-1103201221010012-0210212333003320-3122113112330300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2331123320010023-0031113211011031-3101232031213203-0233113010003232-1003112301031302-1021000133110210-1130022030110312-0233122202122000"></a>

## service_policies_from_namespace — service_policies_from_namespace / 100002031112 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2100233330113013-1011020101321013-1113020011303120-3332020200220223-1130023001010003-0211031020201213-0121011023231220-3020101330032130)
- service_policies_from_namespace

<a id="canonical-1031233021121231-2321310131030302-1012021303323303-3323220310330020-1312131033213310-1332222212230013-3001022313101010-0030103232201320"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

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

<a id="canonical-3312332311300211-1032312333213220-3302211222200200-1001012003210312-3031022313210230-2011003020321002-2000023111013331-2103030021232001"></a>

## Direct properties — service_policies_from_namespace / 100002031112 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2122022010121320-0313112121033303-1202000011130203-1010111232222121-1320310003021201-3221333133310220-2212010110323012-2203221033001212"></a>

## Next pages — service_policies_from_namespace / 100002031112 / 4

- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2100233330113013-1011020101321013-1113020011303120-3332020200220223-1130023001010003-0211031020201213-0121011023231220-3020101330032130)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)

<a id="canonical-2320000221111121-2211321333322221-0131123101031120-1011222001302102-3020022130323223-1330300011100301-1023000101320102-3123222213330010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2100010213001103-1333021202021121-3221103133013202-2302001133001000-0320011002310230-0121223222302102-0332021000313101-1130220032302232"></a>

## udp — udp / 330321302322 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2100233330113013-1011020101321013-1113020011303120-3332020200220223-1130023001010003-0211031020201213-0121011023231220-3020101330032130)
- udp

<a id="canonical-0322233120323103-3131330232221001-0022310300013103-1231002210112131-3023010133201203-2123111033231330-1330110310311130-3111201100222010"></a>

Type: `"single"`. Computed.

Enable this option

Upstream description:

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

<a id="canonical-2023122012032203-3231121020021131-0302212322312122-1310300111213312-0230202000202222-3300101233110311-2001200133220031-1023111213131101"></a>

## Direct properties — udp / 330321302322 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0312232333311213-0202102220321132-3032133232322010-2212333010033301-0020003133032312-2123333232311002-3203302331220332-3121322201023303"></a>

## Next pages — udp / 330321302322 / 4

- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2100233330113013-1011020101321013-1113020011303120-3332020200220223-1130023001010003-0211031020201213-0121011023231220-3020101330032130)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203)
