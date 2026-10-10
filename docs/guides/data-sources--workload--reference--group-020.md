---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-0113223111021000-3232330101122312-2122223231322002-0111100001210313-3303112120020230-1122133031012002-2203013223322310-1011112213223001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_in_cluster.port.info` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_in_cluster](data-sources--workload--reference--group-019.md#canonical-2000012203330031-3013213320003321-1130203203021012-0301212320010201-2013312012132313-2201113011320310-2123010110120120-3030200031122122)
- [stateful_service.advertise_options.advertise_in_cluster.port](data-sources--workload--reference--group-019.md#canonical-0132121001321323-0100130033213110-1210023033132202-1313133133131203-0323330100012311-1222313130120133-3020102231033032-1132100211233330)
- stateful_service.advertise_options.advertise_in_cluster.port.info

<a id="canonical-1322301030011032-3220302120320000-1013330202002000-3332000320121102-0001011121021112-2210211102032021-1112012200030101-3123022111222233"></a>

Type: `"single"`. Computed.

Port Information. Port information.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-target_port_choice": "[\"same_as_port\",\"target_port\"]"
}
```

<a id="canonical-2011222213120011-0111310132133211-3202131121300332-0123200212102122-2132033203332322-2031102200120302-1202213222213022-2003212202020032"></a>

### Direct properties for `stateful_service.advertise_options.advertise_in_cluster.port.info`

<a id="canonical-3221022230133011-0233310332033110-3311113130312113-2220122323012111-1222130022221231-0012030320321200-1003310100213130-0131001113301230"></a>

#### `stateful_service.advertise_options.advertise_in_cluster.port.info.port` property

Type: `"number"`. Computed.

Port. Port the workload can be reached on.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-0030123230000113-0222023013133002-1022132303220132-1123330013011312-0103220211112010-1101231122303220-0020130001011223-1220303203333200"></a>

<a id="canonical-3011300111110230-1210002122323002-0332332333010032-2012123301000331-1130313320033003-1323130111333201-3211320022110101-1112310220301230"></a>

#### `stateful_service.advertise_options.advertise_in_cluster.port.info.protocol` property

Type: `"string"`. Computed.

\[Enum: PROTOCOL\_TCP|PROTOCOL\_HTTP|PROTOCOL\_HTTP2|PROTOCOL\_TLS\_WITH\_SNI|PROTOCOL\_UDP\] Type
of protocol - PROTOCOL\_TCP: TCP TCP - PROTOCOL\_HTTP: HTTP HTTP - PROTOCOL\_HTTP2: HTTP2 HTTP2 -
PROTOCOL\_TLS\_WITH\_SNI: TLS with SNI TLS with SNI - PROTOCOL\_UDP: UDP UDP. Possible values are
\`PROTOCOL\_TCP\`, \`PROTOCOL\_HTTP\`, \`PROTOCOL\_HTTP2\`, \`PROTOCOL\_TLS\_WITH\_SNI\`,
\`PROTOCOL\_UDP\`. Defaults to \`PROTOCOL\_TCP\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "PROTOCOL_TCP",
  "enum": [
    "PROTOCOL_TCP",
    "PROTOCOL_HTTP",
    "PROTOCOL_HTTP2",
    "PROTOCOL_TLS_WITH_SNI",
    "PROTOCOL_UDP"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [same_as_port](data-sources--workload--reference--group-020.md#canonical-1310301010322211-1231321021220312-2211231301331003-0322322213200331-1202200133003200-1031212012023332-1102013300323202-2003021032123232): complete subsection reference.

<a id="canonical-1022013312312201-0023312100332032-0002000031331211-1133202122230320-1311333330012330-3302120023022022-2001020311213302-1112313230122312"></a>

<a id="canonical-3231102020203110-3332202310123232-0203110322002330-3233203000221103-1332201321222221-2133223021311103-3301011013303323-2213321322322000"></a>

#### `stateful_service.advertise_options.advertise_in_cluster.port.info.target_port` property

Type: `"number"`. Computed.

Exclusive with \[same\_as\_port\] Port the workload is listening on.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-1310301010322211-1231321021220312-2211231301331003-0322322213200331-1202200133003200-1031212012023332-1102013300323202-2003021032123232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_in_cluster.port.info.same_as_port` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_in_cluster](data-sources--workload--reference--group-019.md#canonical-2000012203330031-3013213320003321-1130203203021012-0301212320010201-2013312012132313-2201113011320310-2123010110120120-3030200031122122)
- [stateful_service.advertise_options.advertise_in_cluster.port](data-sources--workload--reference--group-019.md#canonical-0132121001321323-0100130033213110-1210023033132202-1313133133131203-0323330100012311-1222313130120133-3020102231033032-1132100211233330)
- [stateful_service.advertise_options.advertise_in_cluster.port.info](data-sources--workload--reference--group-020.md#canonical-0113223111021000-3232330101122312-2122223231322002-0111100001210313-3303112120020230-1122133031012002-2203013223322310-1011112213223001)
- stateful_service.advertise_options.advertise_in_cluster.port.info.same_as_port

<a id="canonical-2022003001321213-3232300332033003-0332232101233122-1032133222101311-1200022012202133-1203233011323230-1203132210131300-0122311233023301"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- stateful_service.advertise_options.advertise_on_public

<a id="canonical-1130231230132023-0311213020112213-1311120122001013-0200233312123132-3321230322121301-0101332122221100-2310123211131011-2223031321103312"></a>

Type: `"single"`. Computed.

Advertise this workload via loadbalancer on internet with default VIP.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-advertise_choice": "[\"multi_ports\",\"port\"]"
}
```

<a id="canonical-0322030311320221-0220011203313021-2030101233013213-0020100223102100-1320011003233023-1012102113122031-1332102112202110-1012230301301213"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public`

- [multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333): complete subsection reference.

- [port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303): complete subsection reference.

<a id="canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.multi_ports` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- stateful_service.advertise_options.advertise_on_public.multi_ports

<a id="canonical-3121300200010120-0123001020033203-1303111311032302-2022010203312223-3203102121003233-0112133321321110-0313232302320333-3023121210303020"></a>

Type: `"single"`. Computed.

Advertise Multiple Ports. Advertise multiple ports.

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

<a id="canonical-1120202331300012-0333313030223201-3012312120203331-1001000232332201-3302002000322120-1302333203312031-1033113021013130-2100131320310121"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.multi_ports`

- [ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110): complete subsection reference.

<a id="canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.multi_ports.ports` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports

<a id="canonical-2310012202032320-2332322033032033-0212213323312131-0202022020233102-0103201121032000-2202200013032221-3123023331300312-2301300230230112"></a>

Type: `"list"`. Computed.

Ports. Ports to advertise.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-1120022331302201-3130101001300112-3213032230220010-2121110112033000-1101020303222310-3030132333233131-2011123023003101-0222323222230333"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.multi_ports.ports`

- [http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230): complete subsection reference.

- [port](data-sources--workload--reference--group-023.md#canonical-2032122332230132-1232301210033133-2310220303021120-1333300211130330-0130020012323310-0231032130121032-0330110103233210-2320011201211330): complete subsection reference.

- [tcp_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0023311210022323-0202111021131210-2023111033133223-1012103322100303-2311011122322021-0003122222031302-1322312130010112-3220023131111022): complete subsection reference.

<a id="canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer

<a id="canonical-1102002201101312-3212223020020102-1331322130132310-3121033213331301-1332111020212103-3332322000211311-3030300002122300-2132231211330312"></a>

Type: `"single"`. Computed.

Configuration parameter for http loadbalancer.

Additional upstream details:

HTTP/HTTPS Load balancer.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-loadbalancer_type": "[\"http\",\"https\",\"https_auto_cert\"]",
  "x-ves-oneof-field-route_choice": "[\"default_route\",\"specific_routes\"]"
}
```

<a id="canonical-1310231331332233-3113230123023333-0011300300303011-0301001230111213-2032201221120002-2021010122323321-2013000231221031-2131021000122321"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer`

- [default_route](data-sources--workload--reference--group-020.md#canonical-2030310202300112-3303300032130110-0103103322303032-0211102122210322-0023002233231012-2302312031202003-1110220232133120-3323320120132322): complete subsection reference.

<a id="canonical-2211100303210131-0130031222133022-3032213313211002-2330130210023102-1103233332222323-2110301201001101-1220231111012312-0222332111231230"></a>

<a id="canonical-0203230020022021-0201303200302201-0101120030233322-0222120102333212-1213212031230310-0322031130023212-1203200120033032-1003230221303333"></a>

#### `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.domains` property

Type: `["list", "string"]`. Computed.

List of domains (host/authority header) that will be matched to loadbalancer. Wildcard hosts are
supported in the suffix or prefix form Domain search order: 1. Exact domain names: \`\` is invalid
Domains are also used for SNI matching if the loadbalancer type is HTTPS Domains also indicate the..

Additional upstream details:

A list of domains (host/authority header) that will be matched to loadbalancer. Exact domain names:
\`\`www&#46;example.com\`\`. &#8203;2. Prefix domain wildcards: \`\`\*.example.com\`\` or
\`\`\*.bar.example.com\`\`. &#8203;3. Special wildcard \`\`\*\`\` matching any domain. Wildcard will
not match empty string. E.g. \`\`\*.example.com\`\` will match \`\`bar.example.com\`\` and
\`\`baz-bar.example.com\`\` but not \`\`.example.com\`\`. The longest wildcards match first.
Wildcards must match a whole DNS label. E.g. \`\`\*.example.com\`\` and \*.bar.example.com are
valid, however \`\`\*bar.example.com\`\` or \`\`\*-bar.example.com\`\` is invalid

Domains are also used for SNI matching if the loadbalancer type is HTTPS Domains also indicate the
list of names for which DNS resolution will be done by VER.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [http](data-sources--workload--reference--group-020.md#canonical-0112310331012222-3310301101023101-3030311303122110-1002120000131211-2132222333033211-0221132323013310-2331200022030010-1001111302103012): complete subsection reference.

- [https](data-sources--workload--reference--group-020.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231): complete subsection reference.

- [https_auto_cert](data-sources--workload--reference--group-021.md#canonical-3120032122030013-2103202012301201-3331001221002131-1221031330101133-3201220221020011-0212213033013002-3102212233210301-0200220030321010): complete subsection reference.

- [specific_routes](data-sources--workload--reference--group-022.md#canonical-0200311133303313-1102333013020100-0231130033310130-0220330111033330-0233120333330111-0310003313220103-3012130211002322-2030012200210113): complete subsection reference.

<a id="canonical-2030310202300112-3303300032130110-0103103322303032-0211102122210322-0023002233231012-2302312031202003-1110220232133120-3323320120132322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route

<a id="canonical-2321021000121222-3213213111123223-1021111301331300-2312222301003103-1230000233131301-1301233122221120-2203123221221202-2023013110203022"></a>

Type: `"single"`. Computed.

Configuration parameter for default route.

Additional upstream details:

Default route matching all APIs.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-host_rewrite_params": "[\"auto_host_rewrite\",\"disable_host_rewrite\",\"host_rewrite\"]"
}
```

<a id="canonical-2003021330332012-3010011320312012-2213033033220302-2330111300130122-0303012223112322-1302021031210031-0111113200320301-2100213233120331"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route`

- [auto_host_rewrite](data-sources--workload--reference--group-020.md#canonical-2312133100231010-1132023310002323-2221211231002103-2232301033231112-1230121030131330-3203001022103000-0012103030331233-3332211021001220): complete subsection reference.

- [disable_host_rewrite](data-sources--workload--reference--group-020.md#canonical-2321022330333131-1333102212121100-3211320232032223-3310311201013132-2233202021220201-1222002100310323-1311203223102201-3100031132310112): complete subsection reference.

<a id="canonical-2233221011211332-1310130333332232-0303132322121203-3233022220222121-2233310100221012-1201033130001010-1201131022013200-2212110010332332"></a>

<a id="canonical-0132203222332321-2110301111010333-2231213001132213-1323002003021220-2020220213023212-3212022320300130-1321013311202230-0232230000321121"></a>

#### `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route.host_rewrite` property

Type: `"string"`. Computed.

Exclusive with \[auto\_host\_rewrite disable\_host\_rewrite\] Host header will be swapped with this
value.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="canonical-2312133100231010-1132023310002323-2221211231002103-2232301033231112-1230121030131330-3203001022103000-0012103030331233-3332211021001220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route.auto_host_rewrite` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route](data-sources--workload--reference--group-020.md#canonical-2030310202300112-3303300032130110-0103103322303032-0211102122210322-0023002233231012-2302312031202003-1110220232133120-3323320120132322)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route.auto_host_rewrite

<a id="canonical-3030110323132322-2200113011321301-1130032122031013-0102303322212233-3112013313203022-2021102322123312-2312102013303100-3022201030320133"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2321022330333131-1333102212121100-3211320232032223-3310311201013132-2233202021220201-1222002100310323-1311203223102201-3100031132310112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route.disable_host_rewrite` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route](data-sources--workload--reference--group-020.md#canonical-2030310202300112-3303300032130110-0103103322303032-0211102122210322-0023002233231012-2302312031202003-1110220232133120-3323320120132322)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route.disable_host_rewrite

<a id="canonical-1332303113222011-0210122330122231-2200302001321222-3133033010320232-3202112122120023-2112213131302133-2323330223131103-2122111230333320"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0112310331012222-3310301101023101-3030311303122110-1002120000131211-2132222333033211-0221132323013310-2331200022030010-1001111302103012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.http` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.http

<a id="canonical-3130301310003333-2033002331303022-3213301223311103-1021123220112320-1330000132102231-1020000220003221-0200222021022330-0300003103130113"></a>

Type: `"single"`. Computed.

HTTP Choice. Choice for selecting HTTP proxy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_choice": "[\"port\",\"port_ranges\"]"
}
```

<a id="canonical-3201221012221123-0120220032220033-3203223301032021-1202121000012111-0313021303130131-3021123030032023-2221220112032132-0311333201132021"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.http`

<a id="canonical-0001303201003022-1130301031232000-3222311021130311-0110321003332331-0222103110230321-2303220002003212-0122322033330300-1001302332100210"></a>

#### `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.http.dns_volterra_managed` property

Type: `"bool"`. Computed.

DNS records for domains will be managed automatically by F5 Distributed Cloud. As a prerequisite,
the domain must be delegated to F5 Distributed Cloud using Delegated domain feature or a DNS CNAME
record should be created in your DNS provider's portal.

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

<a id="canonical-2323121110323320-1222201331003321-2120310201323031-2222311322313201-3203303102132220-3320203003103121-2002021101123312-1332322202231113"></a>

<a id="canonical-0330221330123120-1323300100133001-1222012011103201-0002310303221301-2030021220200011-2221100232302232-2021311102301203-2311122232323211"></a>

#### `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.http.port` property

Type: `"number"`. Computed.

Exclusive with \[port\_ranges\] HTTP port to Listen.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-1033322120122221-3330132312131032-0311131011210211-3103000001313312-0003232211100331-2130223132320223-0211323323112022-0230210322110100"></a>

<a id="canonical-0002310320302101-3332122323320132-3030211200311133-3332002232203131-2323322002103233-2010313232323201-2120020212010021-0123320320232230"></a>

#### `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.http.port_ranges` property

Type: `"string"`. Computed.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Additional upstream details:

Each port range consists of a single port or two ports separated by "-".

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https

<a id="canonical-1123030011002223-0102231300311003-2121230032211023-0010032001111100-3120331310120332-1312322120021013-2321310103223130-1121000223210103"></a>

Type: `"single"`. Computed.

Choice for selecting HTTP proxy with bring your own certificates.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-default_lb_choice": "[\"default_loadbalancer\",\"non_default_loadbalancer\"]",
  "x-ves-oneof-field-path_normalize_choice": "[\"disable_path_normalize\",\"enable_path_normalize\"]",
  "x-ves-oneof-field-port_choice": "[\"port\",\"port_ranges\"]",
  "x-ves-oneof-field-server_header_choice": "[\"append_server_name\",\"default_header\",\"pass_through\",\"server_name\"]",
  "x-ves-oneof-field-tls_certificates_choice": "[\"tls_cert_params\",\"tls_parameters\"]"
}
```

<a id="canonical-0320120133301203-2230032020103133-2132220230231000-0020232102322310-2223220220333110-2103010311130311-1103233222221001-2212222133203302"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https`

<a id="canonical-2311233231201300-1000303020202212-1023202130000221-0113122023120123-3022102123012130-0311102321303113-1033010321331330-1301312032212211"></a>

#### `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.add_hsts` property

Type: `"bool"`. Computed.

Add HTTP Strict-Transport-Security response header.

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

<a id="canonical-0202301330000230-0312333130312320-0023233220303302-1210001011030133-3121203132222132-1033231010003121-2131023010001310-0100030333313023"></a>

<a id="canonical-2230200031003011-1230032031223211-0031001332223212-3201031202311030-2200030311333102-2000003021003021-3202333330321212-2032323002000130"></a>

#### `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.append_server_name` property

Type: `"string"`. Computed.

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

- [coalescing_options](data-sources--workload--reference--group-020.md#canonical-0223013021320231-2210210321322111-3013223300203101-1323100321101202-1032021203311322-2030033033323022-1130010310100330-0302011323001003): complete subsection reference.

<a id="canonical-2010031033203322-0211222210300130-2321111010300113-2210121100321121-3021133122001111-1303103210200300-1121022231030103-3223221112012303"></a>

<a id="canonical-3100031133033322-0210230110121312-3102130331302222-1203001030231302-3021123213113232-3222200111213003-3312232310301313-3123230301201210"></a>

#### `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.connection_idle_timeout` property

Type: `"number"`. Computed.

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed. Note
that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive. This is
specified in milliseconds. The default value is 2 minutes.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

- [default_header](data-sources--workload--reference--group-020.md#canonical-3122021022021130-2031201222023001-0203131103201023-1223221202213312-3323012032021310-3300120210220111-3321210123213333-2303131033232212): complete subsection reference.

- [default_loadbalancer](data-sources--workload--reference--group-020.md#canonical-3312132213003110-3220223213233203-1013321120233203-0131000211303220-2233231131201013-1011202203021310-3331211003001131-2303221001021310): complete subsection reference.

- [disable_path_normalize](data-sources--workload--reference--group-020.md#canonical-0311110212032330-3222120112023023-3321032333013030-3103101303022011-0120222102333201-3200113323301112-2303001031323212-2120120221322321): complete subsection reference.

- [enable_path_normalize](data-sources--workload--reference--group-020.md#canonical-0230203033333220-1103323303221313-1132312203222022-3311332312300330-0122230203313031-2330333031222120-0132110230101001-3022030003101312): complete subsection reference.

- [http_protocol_options](data-sources--workload--reference--group-020.md#canonical-0331311332303112-3223030300211302-2031331031223130-2022220203222033-0230003311302230-0311213313002123-3230132212113331-2232020110131310): complete subsection reference.

<a id="canonical-0033330330222110-3132231321331320-2123321111020322-1113110032010320-0023301333221013-2212120121002303-0030121013323330-3230303011133100"></a>

<a id="canonical-0301103231103100-0310022111010112-3013110312330113-3020132021100310-1110020011301102-1010201032110032-0010200313012330-2102130031102133"></a>

#### `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_redirect` property

Type: `"bool"`. Computed.

HTTP Redirect to HTTPS. Redirect HTTP traffic to HTTPS.

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

- [non_default_loadbalancer](data-sources--workload--reference--group-020.md#canonical-0011033233320121-3333020002121100-1201121003022320-2303002220331022-3023331020321233-0320223300120033-3021203022022003-2131002203233322): complete subsection reference.

- [pass_through](data-sources--workload--reference--group-020.md#canonical-0000223200231120-3103130233313330-2132331313003321-2203211112112303-0302010122020301-3031301310323333-3132033122212331-3022302332313302): complete subsection reference.

<a id="canonical-1222321232300231-1120312023030321-1333101233313322-1310231303130313-3312202311023032-0102221133121303-1123010110333230-1121203012220203"></a>

<a id="canonical-1200102110120221-0023130221310211-2233100121223132-2000222302121331-0002200203212202-0221123133113230-2122201122132220-2301120200303130"></a>

#### `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.port` property

Type: `"number"`. Computed.

Exclusive with \[port\_ranges\] HTTPS port to Listen.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-3001010111013311-1103303132000023-3210012122013022-1011331010012130-1100302123301331-2113021331031310-0333120030110320-2301213133313322"></a>

<a id="canonical-2131033011030031-0220130121210320-1133033200223000-0133222122323000-0333303103230320-2321023210102322-3301200012031333-2032000011330101"></a>

#### `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.port_ranges` property

Type: `"string"`. Computed.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Additional upstream details:

Each port range consists of a single port or two ports separated by "-".

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0313312001100203-3131001232213013-0032322102321022-0112233332232300-0133121030101311-3100030211311322-3220303223322231-3310301332130222"></a>

<a id="canonical-3302121120102003-1032223112120312-0010110021100221-0232300313033313-2200123222012102-0220331032110021-1220010300313133-3003111213303003"></a>

#### `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.server_name` property

Type: `"string"`. Computed.

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

- [tls_cert_params](data-sources--workload--reference--group-020.md#canonical-0020221002300310-3320210022331231-2110020220012113-3303311310321122-1023231100033031-3133012031200203-1002323202122123-3323330220103313): complete subsection reference.

- [tls_parameters](data-sources--workload--reference--group-020.md#canonical-0010123223300323-3122033232120131-1033211323302101-0202012323131002-0013101320130212-1220113010330021-1032021212320311-0300031310102303): complete subsection reference.

<a id="canonical-0223013021320231-2210210321322111-3013223300203101-1323100321101202-1032021203311322-2030033033323022-1130010310100330-0302011323001003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-020.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options

<a id="canonical-1203020012302030-0001013011021012-0113002313023321-2332002311302210-2313233303021300-3203211012313303-0013312310011230-0221332120322010"></a>

Type: `"single"`. Computed.

TLS connection coalescing configuration (not compatible with mTLS).

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-coalescing_choice": "[\"default_coalescing\",\"strict_coalescing\"]"
}
```

<a id="canonical-2200132110103332-2131320210110001-0010321111113100-3200010021222331-1030211221032131-2303322212101130-1112203321231212-2120021200322130"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options`

- [default_coalescing](data-sources--workload--reference--group-020.md#canonical-1222213323120211-3133013131321230-3103230031310110-1230220133202132-1121323100033323-1022200022001013-3231132033010103-0330112230023020): complete subsection reference.

- [strict_coalescing](data-sources--workload--reference--group-020.md#canonical-1021011213302123-2123303232103013-3021200330320233-1031313321012003-2310000001130323-3230012231310302-1032123110132323-0131220032121320): complete subsection reference.

<a id="canonical-1222213323120211-3133013131321230-3103230031310110-1230220133202132-1121323100033323-1022200022001013-3231132033010103-0330112230023020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options.default_coalescing` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-020.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options](data-sources--workload--reference--group-020.md#canonical-0223013021320231-2210210321322111-3013223300203101-1323100321101202-1032021203311322-2030033033323022-1130010310100330-0302011323001003)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options.default_coalescing

<a id="canonical-3111201011232022-3223222131210302-1112212313230333-3201031122003313-1220213313001203-3330202203312301-1210133321132220-2321201210302213"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default coalescing.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1021011213302123-2123303232103013-3021200330320233-1031313321012003-2310000001130323-3230012231310302-1032123110132323-0131220032121320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options.strict_coalescing` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-020.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options](data-sources--workload--reference--group-020.md#canonical-0223013021320231-2210210321322111-3013223300203101-1323100321101202-1032021203311322-2030033033323022-1130010310100330-0302011323001003)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options.strict_coalescing

<a id="canonical-3110202200031103-0303321121121022-1300021000113233-1323203020211231-2230010121111002-2103330120113321-1111222121111210-1333003100132202"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for strict coalescing.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3122021022021130-2031201222023001-0203131103201023-1223221202213312-3323012032021310-3300120210220111-3321210123213333-2303131033232212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.default_header` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-020.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.default_header

<a id="canonical-2311202111003331-2112022010232200-3230203111332032-2033101203233231-2011322133230323-0033323033310120-3103020312013303-2102221202121310"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default header.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3312132213003110-3220223213233203-1013321120233203-0131000211303220-2233231131201013-1011202203021310-3331211003001131-2303221001021310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.default_loadbalancer` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-020.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.default_loadbalancer

<a id="canonical-2011020033212012-3310030032302102-2031213101201230-3023132322023313-0323203011232230-1012001201330210-2311232333211201-3301100212110211"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default loadbalancer.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0311110212032330-3222120112023023-3321032333013030-3103101303022011-0120222102333201-3200113323301112-2303001031323212-2120120221322321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.disable_path_normalize` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-020.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.disable_path_normalize

<a id="canonical-3113201230223310-1101010202313122-1010010230221211-2312020310013210-3322331001110033-1330220320021001-2300333002031121-3011202101331023"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0230203033333220-1103323303221313-1132312203222022-3311332312300330-0122230203313031-2330333031222120-0132110230101001-3022030003101312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.enable_path_normalize` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-020.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.enable_path_normalize

<a id="canonical-1321333120211302-1323031332311223-3230210330131210-2131300111031201-0123101302311020-1123131301023012-0212133113010013-2330022132002312"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0331311332303112-3223030300211302-2031331031223130-2022220203222033-0230003311302230-0311213313002123-3230132212113331-2232020110131310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-020.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options

<a id="canonical-2233212223322132-2203202033301320-2300130223012101-2230122312222321-3031002102022001-1213010020121300-3303113230232320-0133031200021022"></a>

Type: `"single"`. Computed.

HTTP protocol configuration OPTIONS for downstream connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-http_protocol_choice": "[\"http_protocol_enable_v1_only\",\"http_protocol_enable_v1_v2\",\"http_protocol_enable_v2_only\"]"
}
```

<a id="canonical-3220310033003200-1312203231313011-1132213301013021-0103013213313301-1023103113320000-0123110001101011-2022022233202313-2331220213030322"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options`

- [http_protocol_enable_v1_only](data-sources--workload--reference--group-020.md#canonical-1033131203023032-2102001223320103-0323131200122322-2301231132233112-1223111332323113-3231022031110001-1223033201000032-3310113300131012): complete subsection reference.

- [http_protocol_enable_v1_v2](data-sources--workload--reference--group-020.md#canonical-2203131003213122-3331030200230003-1001023323300301-1023203312002120-2112303110321012-2131131303320221-3032033312210013-2023033121200003): complete subsection reference.

- [http_protocol_enable_v2_only](data-sources--workload--reference--group-020.md#canonical-2330131301031201-1210032120330101-1321113003123120-0103221211320010-1333112300210212-1220220331310110-2030120112022001-2103332101212322): complete subsection reference.

<a id="canonical-1033131203023032-2102001223320103-0323131200122322-2301231132233112-1223111332323113-3231022031110001-1223033201000032-3310113300131012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-020.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-020.md#canonical-0331311332303112-3223030300211302-2031331031223130-2022220203222033-0230003311302230-0311213313002123-3230132212113331-2232020110131310)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only

<a id="canonical-1122311230210313-3100223113201021-1210201003033133-1211220033201102-0233122111221133-0131130333103323-0001120332221133-0013222201103333"></a>

Type: `"single"`. Computed.

HTTP/1.1 Protocol OPTIONS for downstream connections.

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

<a id="canonical-1110312012230003-1023123011112210-3120220302132332-2022013220101310-1302033302300002-0103112011033103-3311223031102300-3221302210101222"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only`

- [header_transformation](data-sources--workload--reference--group-020.md#canonical-1000301122202312-3011230331001222-1103133100213100-0100322330231321-3210311100100323-1222032203300032-0300300200303201-1102223101130200): complete subsection reference.

<a id="canonical-1000301122202312-3011230331001222-1103133100213100-0100322330231321-3210311100100323-1222032203300032-0300300200303201-1102223101130200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-020.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-020.md#canonical-0331311332303112-3223030300211302-2031331031223130-2022220203222033-0230003311302230-0311213313002123-3230132212113331-2232020110131310)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-020.md#canonical-1033131203023032-2102001223320103-0323131200122322-2301231132233112-1223111332323113-3231022031110001-1223033201000032-3310113300131012)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="canonical-2130030310302323-1213012112303232-1100121133013132-2001131232012211-0011012131222110-1323231303133030-3311302312001121-3302232131102000"></a>

Type: `"single"`. Computed.

Header Transformation OPTIONS for HTTP/1.1 request/response headers.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-header_transformation_choice": "[\"default_header_transformation\",\"preserve_case_header_transformation\",\"proper_case_header_transformation\"]"
}
```

<a id="canonical-3300213131233303-1200332103201230-0320333300010331-0323021021202133-3100000300211002-0122021132110000-0120302002320022-1333130011310310"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation`

- [default_header_transformation](data-sources--workload--reference--group-020.md#canonical-2013002213000101-1330333123203131-3300111133110113-3333110213211322-0013302210300232-3202020310201112-3001233321300232-1100302301212021): complete subsection reference.

- [preserve_case_header_transformation](data-sources--workload--reference--group-020.md#canonical-1200211130220221-3133230021110123-3003231020332330-0213130301203020-2300203200121212-0301202202201310-3321211010321203-2030023311211121): complete subsection reference.

- [proper_case_header_transformation](data-sources--workload--reference--group-020.md#canonical-2112201210020123-2221033111203013-3202202012112111-0123232323311222-0013103312310011-1103112023303220-0201102213220121-0220331221020210): complete subsection reference.

<a id="canonical-2013002213000101-1330333123203131-3300111133110113-3333110213211322-0013302210300232-3202020310201112-3001233321300232-1100302301212021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-020.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-020.md#canonical-0331311332303112-3223030300211302-2031331031223130-2022220203222033-0230003311302230-0311213313002123-3230132212113331-2232020110131310)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-020.md#canonical-1033131203023032-2102001223320103-0323131200122322-2301231132233112-1223111332323113-3231022031110001-1223033201000032-3310113300131012)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-020.md#canonical-1000301122202312-3011230331001222-1103133100213100-0100322330231321-3210311100100323-1222032203300032-0300300200303201-1102223101130200)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

<a id="canonical-1032131320333323-2211001000012001-2322031113022232-1200033201022102-0211133031002332-0100301201131202-3020031012223100-3031032202212222"></a>

Type: `["object", {}]`. Computed.

Use the platform's current default HTTP header transformation behavior.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1200211130220221-3133230021110123-3003231020332330-0213130301203020-2300203200121212-0301202202201310-3321211010321203-2030023311211121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-020.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-020.md#canonical-0331311332303112-3223030300211302-2031331031223130-2022220203222033-0230003311302230-0311213313002123-3230132212113331-2232020110131310)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-020.md#canonical-1033131203023032-2102001223320103-0323131200122322-2301231132233112-1223111332323113-3231022031110001-1223033201000032-3310113300131012)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-020.md#canonical-1000301122202312-3011230331001222-1103133100213100-0100322330231321-3210311100100323-1222032203300032-0300300200303201-1102223101130200)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation

<a id="canonical-2113302011011022-1302022023202000-0220201130022221-0332133302203212-2221121101133330-3121333220130211-0003003002003130-3321112111201102"></a>

Type: `["object", {}]`. Computed.

Preserve HTTP header-name case when upstream case must remain unchanged.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2112201210020123-2221033111203013-3202202012112111-0123232323311222-0013103312310011-1103112023303220-0201102213220121-0220331221020210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-020.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-020.md#canonical-0331311332303112-3223030300211302-2031331031223130-2022220203222033-0230003311302230-0311213313002123-3230132212113331-2232020110131310)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-020.md#canonical-1033131203023032-2102001223320103-0323131200122322-2301231132233112-1223111332323113-3231022031110001-1223033201000032-3310113300131012)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-020.md#canonical-1000301122202312-3011230331001222-1103133100213100-0100322330231321-3210311100100323-1222032203300032-0300300200303201-1102223101130200)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation

<a id="canonical-2020100112310122-1313220102020333-2320311020321232-0203222031301312-2122103321322003-0011130212120212-0213020330102101-1330313001122220"></a>

Type: `["object", {}]`. Computed.

Transform HTTP header names to proper case when explicit transformation is required.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2203131003213122-3331030200230003-1001023323300301-1023203312002120-2112303110321012-2131131303320221-3032033312210013-2023033121200003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_v2` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-020.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-020.md#canonical-0331311332303112-3223030300211302-2031331031223130-2022220203222033-0230003311302230-0311213313002123-3230132212113331-2232020110131310)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_v2

<a id="canonical-2331123030131123-2002202100033231-3322101322212130-3001333300210022-3212212320232230-0101120131132031-0010121332233301-0232112220233011"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for http protocol enable v1 v2.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2330131301031201-1210032120330101-1321113003123120-0103221211320010-1333112300210212-1220220331310110-2030120112022001-2103332101212322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v2_only` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-020.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-020.md#canonical-0331311332303112-3223030300211302-2031331031223130-2022220203222033-0230003311302230-0311213313002123-3230132212113331-2232020110131310)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v2_only

<a id="canonical-0113333003323322-3213003110033122-2211120330332332-3200310200300310-1302203112202231-0330213030021010-2201222033322133-2131323232222300"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for http protocol enable v2 only.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0011033233320121-3333020002121100-1201121003022320-2303002220331022-3023331020321233-0320223300120033-3021203022022003-2131002203233322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.non_default_loadbalancer` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-020.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.non_default_loadbalancer

<a id="canonical-3302103220201211-1221111323001110-3311302020010322-3023003202032332-0232113322030010-0233310112030010-3332112123210311-0322133323022312"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for non default loadbalancer.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0000223200231120-3103130233313330-2132331313003321-2203211112112303-0302010122020301-3031301310323333-3132033122212331-3022302332313302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.pass_through` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-020.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.pass_through

<a id="canonical-3213302221223101-2120012310120030-2223223110111232-3231113001113103-2300122333120032-0223232332333213-1312320130202303-0303211100222222"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for pass through.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0020221002300310-3320210022331231-2110020220012113-3303311310321122-1023231100033031-3133012031200203-1002323202122123-3323330220103313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-020.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params

<a id="canonical-2302013020310322-1011222101323331-2133212202023023-0011312332233003-0321233203213330-1230232222331121-1310022311310313-3321133020330111"></a>

Type: `"single"`. Computed.

Configuration parameter for tls cert params.

Additional upstream details:

Select TLS Parameters and Certificates.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]"
}
```

<a id="canonical-0211110212300031-3122200012221221-1011331123013100-0333012212210231-2113131312221122-3113110320112210-2123213130321300-1321212233330203"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params`

- [certificates](data-sources--workload--reference--group-020.md#canonical-3020211132312003-2132320021113331-2130210102202222-3121203032130210-2022321103332203-2002131130120123-2232003211133002-1130333010212001): complete subsection reference.

- [no_mtls](data-sources--workload--reference--group-020.md#canonical-3223131212211020-2200001022101101-1300033222302333-0031122200112122-1201200332201032-2331221210210122-1231313302012312-1010003303232203): complete subsection reference.

- [tls_config](data-sources--workload--reference--group-020.md#canonical-3300211012000000-2320102320120012-3300133230232310-0330102233032323-0223002013231003-3202131212320332-3010000213302232-0330030232232230): complete subsection reference.

- [use_mtls](data-sources--workload--reference--group-020.md#canonical-2120202120230123-1300020120013231-2032223101233223-1023310302320101-2320320032132221-0031023101132320-2230232211332312-0332131011020102): complete subsection reference.

<a id="canonical-3020211132312003-2132320021113331-2130210102202222-3121203032130210-2022321103332203-2002131130120123-2232003211133002-1130333010212001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.certificates` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-020.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-020.md#canonical-0020221002300310-3320210022331231-2110020220012113-3303311310321122-1023231100033031-3133012031200203-1002323202122123-3323330220103313)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.certificates

<a id="canonical-1033231102230031-3003330321122313-0032323110013313-0223301033000100-1131231021213111-1113132332021301-0312113032113232-0331221113102102"></a>

Type: `"list"`. Computed.

Select one or more certificates with any domain names.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
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
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1200131303012003-2322102330130022-3203333230323120-1220102000201331-2102010003330200-3031110022230321-2302331012221022-3211321303132102"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.certificates`

<a id="canonical-0000220331230122-2312302022033102-2031031301233320-0323013012100002-1012303130102202-1321102203033300-2110303103313021-2001003111311220"></a>

#### `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.certificates.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1222302220330333-2212130002202210-2203122312332210-1231230300022222-3032230232030102-3322311033232031-2022022300000102-1322213132331221"></a>

<a id="canonical-0103002132203100-0323112121022230-0230130311103222-2120112100231231-1330012121122323-2110121211223210-3031201311010123-1100120120013031"></a>

#### `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.certificates.namespace` property

Type: `"string"`. Computed.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-0012333213213201-0321312113333121-1023313003002310-2233313110001221-1123103033322310-2102101011322033-3113210332310123-3223130313212123"></a>

<a id="canonical-2230210201013012-1211120123221210-1333302332313320-1323300333111201-3312230110223020-1122111213110023-0311312121322202-0101302211021012"></a>

#### `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.certificates.tenant` property

Type: `"string"`. Computed.

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3223131212211020-2200001022101101-1300033222302333-0031122200112122-1201200332201032-2331221210210122-1231313302012312-1010003303232203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.no_mtls` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-020.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-020.md#canonical-0020221002300310-3320210022331231-2110020220012113-3303311310321122-1023231100033031-3133012031200203-1002323202122123-3323330220103313)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.no_mtls

<a id="canonical-2201330023113121-0103131303302033-0120230001023332-1001111323030113-1230033312212301-0320112200022312-2033201132020320-3101313130323323"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3300211012000000-2320102320120012-3300133230232310-0330102233032323-0223002013231003-3202131212320332-3010000213302232-0330030232232230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-020.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-020.md#canonical-0020221002300310-3320210022331231-2110020220012113-3303311310321122-1023231100033031-3133012031200203-1002323202122123-3323330220103313)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config

<a id="canonical-1103202111030200-0102212233223010-0300211020230312-2313102332011332-1301222211131221-3131321002122132-3330201331203002-3112313001202233"></a>

Type: `"single"`. Computed.

This defines various OPTIONS to configure TLS configuration parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

<a id="canonical-0001113023233320-1102332213302033-2300022120131023-3130033312223122-2012311100320100-3233020002123303-3301231132003100-1103211011201312"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config`

- [custom_security](data-sources--workload--reference--group-020.md#canonical-1202101111222012-3002332031013121-2120032333220101-0100133001331210-2222120200120130-0131113102130032-2203223223311223-0300113121310002): complete subsection reference.

- [default_security](data-sources--workload--reference--group-020.md#canonical-3222102130230311-0231110001321300-0011230131323031-1020131303012223-0130201332100213-2032113133300312-0322320210210001-3332232002211021): complete subsection reference.

- [low_security](data-sources--workload--reference--group-020.md#canonical-1312002002313030-0033001230133003-3003333310321212-2201302000122233-2221200310113302-2032001032230302-1210113010131121-0022032131222010): complete subsection reference.

- [medium_security](data-sources--workload--reference--group-020.md#canonical-0112110021323100-0222131201310330-2023013210131131-2303321032320031-3233033313023032-2101013020213200-2300300333133001-0202101221002022): complete subsection reference.

<a id="canonical-1202101111222012-3002332031013121-2120032333220101-0100133001331210-2222120200120130-0131113102130032-2203223223311223-0300113121310002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.custom_security` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-020.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-020.md#canonical-0020221002300310-3320210022331231-2110020220012113-3303311310321122-1023231100033031-3133012031200203-1002323202122123-3323330220103313)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config](data-sources--workload--reference--group-020.md#canonical-3300211012000000-2320102320120012-3300133230232310-0330102233032323-0223002013231003-3202131212320332-3010000213302232-0330030232232230)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.custom_security

<a id="canonical-3121001020131300-0230230322303010-2020221332211102-3000110300200020-3203112201202132-0122332231033010-1020122201112100-2111331030302011"></a>

Type: `"single"`. Computed.

This defines TLS protocol config including min/max versions and allowed ciphers.

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

<a id="canonical-0022012213000013-2313022232000212-3302032133011321-0110031031012011-1123200210203031-3231023131212100-2201333320231102-3022003012303111"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.custom_security`

<a id="canonical-2322303121210022-0100231310323201-0031332112310013-1130312120030000-3010312013120300-3333133012022132-1300113103113023-1300021300032103"></a>

#### `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.custom_security.cipher_suites` property

Type: `["list", "string"]`. Computed.

The TLS listener will only support the specified cipher list.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
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
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0013233223301130-0010313130323002-2102332332010231-3023221101300113-2311210122331131-0120302011001103-0000230321000011-3103222201313211"></a>

<a id="canonical-1011100021313311-3310213310213013-1310111303232303-1332012302031323-3033110123323020-0022000201101221-1211202010101223-3201212302211120"></a>

#### `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.custom_security.max_version` property

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1212221230103111-3131212211021302-3312120001213010-0333331201100303-2101013030233131-1020320010100221-2223133302102323-1323311320122200"></a>

<a id="canonical-3302320030113211-2003231133030223-3011133322310222-1032221233010210-2130131113310113-1131220222130031-3132011123121123-0321103233112322"></a>

#### `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.custom_security.min_version` property

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3222102130230311-0231110001321300-0011230131323031-1020131303012223-0130201332100213-2032113133300312-0322320210210001-3332232002211021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.default_security` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-020.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-020.md#canonical-0020221002300310-3320210022331231-2110020220012113-3303311310321122-1023231100033031-3133012031200203-1002323202122123-3323330220103313)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config](data-sources--workload--reference--group-020.md#canonical-3300211012000000-2320102320120012-3300133230232310-0330102233032323-0223002013231003-3202131212320332-3010000213302232-0330030232232230)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.default_security

<a id="canonical-2333320333032221-3033301301213332-0322222220333022-1110332031230032-3123203203332302-0331113130112130-2033201011010301-0310211111003330"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1312002002313030-0033001230133003-3003333310321212-2201302000122233-2221200310113302-2032001032230302-1210113010131121-0022032131222010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.low_security` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-020.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-020.md#canonical-0020221002300310-3320210022331231-2110020220012113-3303311310321122-1023231100033031-3133012031200203-1002323202122123-3323330220103313)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config](data-sources--workload--reference--group-020.md#canonical-3300211012000000-2320102320120012-3300133230232310-0330102233032323-0223002013231003-3202131212320332-3010000213302232-0330030232232230)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.low_security

<a id="canonical-0000022131332011-3111330210000222-1030113013121210-3121200212111022-0311211132213113-1222232321312201-0120021132010312-0100120122131112"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0112110021323100-0222131201310330-2023013210131131-2303321032320031-3233033313023032-2101013020213200-2300300333133001-0202101221002022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.medium_security` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-020.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-020.md#canonical-0020221002300310-3320210022331231-2110020220012113-3303311310321122-1023231100033031-3133012031200203-1002323202122123-3323330220103313)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config](data-sources--workload--reference--group-020.md#canonical-3300211012000000-2320102320120012-3300133230232310-0330102233032323-0223002013231003-3202131212320332-3010000213302232-0330030232232230)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.medium_security

<a id="canonical-3312113202131202-2212212223010303-1100020011131111-0121302321311213-2103012000312313-0201322231320020-2130230332130202-1220320203203011"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2120202120230123-1300020120013231-2032223101233223-1023310302320101-2320320032132221-0031023101132320-2230232211332312-0332131011020102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-020.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-020.md#canonical-0020221002300310-3320210022331231-2110020220012113-3303311310321122-1023231100033031-3133012031200203-1002323202122123-3323330220103313)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls

<a id="canonical-2310330323202330-3013131313331233-3121002021111310-0022332221101111-0031013110301211-2232221323101330-0302320033102321-3232223300123133"></a>

Type: `"single"`. Computed.

Validation context for downstream client TLS connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-crl_choice": "[\"crl\",\"no_crl\"]",
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]",
  "x-ves-oneof-field-xfcc_header": "[\"xfcc_disabled\",\"xfcc_options\"]"
}
```

<a id="canonical-1331211102131330-3312133313312202-2321033103120323-3303223310200022-0211331232000221-0123232120323320-3112212002122220-1230313313022312"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls`

<a id="canonical-3321210203200101-3122303011023333-1003223102331212-2112303233101001-0030220233322201-3110313000131213-2203133131000323-1213032120203320"></a>

#### `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.client_certificate_optional` property

Type: `"bool"`. Computed.

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated. If the client
does not provide a certificate, the connection will be accepted.

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

- [crl](data-sources--workload--reference--group-020.md#canonical-1310213332120302-0103322220333223-1233111332130300-0311100132002131-0311133030130313-3311110112032331-0033021231223221-3320330103113330): complete subsection reference.

- [no_crl](data-sources--workload--reference--group-020.md#canonical-3330111311003230-1013210213213010-3302003001221223-2022033011323212-0322102310001103-1012310011203120-1023110220010313-2110020233313100): complete subsection reference.

- [trusted_ca](data-sources--workload--reference--group-020.md#canonical-2211301003212230-3232001131031222-1123233020020033-3303312111320132-0110303113321131-1210032210322331-3120003200201311-0221322022323101): complete subsection reference.

<a id="canonical-3013231120002001-0212101031210220-0113233212130212-0221112321121321-2002213133003201-1331302230023112-1330113202031201-3221010102303132"></a>

<a id="canonical-2312320333231330-0201313322020033-1101302110013200-0332300113020323-1132223030103232-2212230320212320-2022122203110100-1110121031330022"></a>

#### `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca_url` property

Type: `"string"`. Computed.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

- [xfcc_disabled](data-sources--workload--reference--group-020.md#canonical-0030111210231013-2222220121122130-2010310023010310-1130132313031002-1330222301322013-1210011310321303-2232121032200102-2310010213332002): complete subsection reference.

- [xfcc_options](data-sources--workload--reference--group-020.md#canonical-3221100202303211-1330113112311332-1132323132323111-2020203132010123-0100312032302132-3010031232110212-2323021103032020-2102132331230010): complete subsection reference.

<a id="canonical-1310213332120302-0103322220333223-1233111332130300-0311100132002131-0311133030130313-3311110112032331-0033021231223221-3320330103113330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.crl` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-020.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-020.md#canonical-0020221002300310-3320210022331231-2110020220012113-3303311310321122-1023231100033031-3133012031200203-1002323202122123-3323330220103313)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-020.md#canonical-2120202120230123-1300020120013231-2032223101233223-1023310302320101-2320320032132221-0031023101132320-2230232211332312-0332131011020102)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.crl

<a id="canonical-0132112022122020-3331210302212100-1112313002323020-1312321303210211-2330302231012123-3132110222202012-3231030113023131-2201030210011033"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3031230321332233-1122022111302021-0213221110031230-3132303002000030-2113223112323223-1002033133002123-2220132321330323-1023021211021100"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.crl`

<a id="canonical-1002330032111230-1303130211231033-3212132103312003-2302033302130022-3332332231201332-0013022312321031-3000000101222322-3312001123021110"></a>

#### `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.crl.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0221011030200332-2122132330003310-3332200020023310-0202022002233131-1013202220131320-3321203312001300-2012231222001331-1332031132122011"></a>

<a id="canonical-1330102303100012-0301012111322231-2232111110221203-2222020113211321-0330213323023212-2321231111202233-1010333113110223-1031213131203113"></a>

#### `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.crl.namespace` property

Type: `"string"`. Computed.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1131320230231002-1021002133130230-1001101303030320-2030003221121023-0113100220223303-3120010203233232-3022311202322233-0131332233202023"></a>

<a id="canonical-2120222023230301-2131320123000330-0210111213030213-2333303012110232-0100022221201201-1201022022331233-3330202222303100-2120132130230321"></a>

#### `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.crl.tenant` property

Type: `"string"`. Computed.

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3330111311003230-1013210213213010-3302003001221223-2022033011323212-0322102310001103-1012310011203120-1023110220010313-2110020233313100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.no_crl` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-020.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-020.md#canonical-0020221002300310-3320210022331231-2110020220012113-3303311310321122-1023231100033031-3133012031200203-1002323202122123-3323330220103313)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-020.md#canonical-2120202120230123-1300020120013231-2032223101233223-1023310302320101-2320320032132221-0031023101132320-2230232211332312-0332131011020102)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.no_crl

<a id="canonical-3322331220000112-1112303031111110-3102132101120030-0310210322033131-0130002020103201-2030320223110322-3031303132113103-2322211123020112"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2211301003212230-3232001131031222-1123233020020033-3303312111320132-0110303113321131-1210032210322331-3120003200201311-0221322022323101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-020.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-020.md#canonical-0020221002300310-3320210022331231-2110020220012113-3303311310321122-1023231100033031-3133012031200203-1002323202122123-3323330220103313)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-020.md#canonical-2120202120230123-1300020120013231-2032223101233223-1023310302320101-2320320032132221-0031023101132320-2230232211332312-0332131011020102)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca

<a id="canonical-1100232020213003-3312331302332222-2032320333311030-0232020232122203-2101331223200120-0223231031102022-2222013011222212-0303222332331232"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0231211300213112-3303300120132233-1023021303230123-2111102132122002-1322232030213122-0332231202130212-1010231122212313-1011210120123232"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca`

<a id="canonical-3021002003302121-0333220030022021-1331221022212331-2200003232010331-2230013200220010-1102121201002221-1312100033231321-1303330223032021"></a>

#### `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3011112333030103-2332320102300210-3131033301231220-0321123002121311-0230222313113213-1332203002121003-2330011010231120-0330102212030321"></a>

<a id="canonical-2132102200022323-3322303103310332-0302312100203331-0210320132110112-1103112310102032-1131332223302311-2112230221321000-3121313132111220"></a>

#### `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca.namespace` property

Type: `"string"`. Computed.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-0320131211121030-0231301311123301-2321121201301031-2103332310203211-2002131233310110-0113233130313203-0322113013201111-3202321211133003"></a>

<a id="canonical-1120130012133133-2101021322311300-0101101121010101-1001320113230210-1232022202311030-0132323302231333-0000221113120302-0023303002313000"></a>

#### `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca.tenant` property

Type: `"string"`. Computed.

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-0030111210231013-2222220121122130-2010310023010310-1130132313031002-1330222301322013-1210011310321303-2232121032200102-2310010213332002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_disabled` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-020.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-020.md#canonical-0020221002300310-3320210022331231-2110020220012113-3303311310321122-1023231100033031-3133012031200203-1002323202122123-3323330220103313)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-020.md#canonical-2120202120230123-1300020120013231-2032223101233223-1023310302320101-2320320032132221-0031023101132320-2230232211332312-0332131011020102)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_disabled

<a id="canonical-3203312331230212-0022231112012232-0201200221212231-0211111031100113-1020313300212121-3231110213333133-1003132331212301-0331122322200312"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3221100202303211-1330113112311332-1132323132323111-2020203132010123-0100312032302132-3010031232110212-2323021103032020-2102132331230010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-020.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-020.md#canonical-0020221002300310-3320210022331231-2110020220012113-3303311310321122-1023231100033031-3133012031200203-1002323202122123-3323330220103313)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-020.md#canonical-2120202120230123-1300020120013231-2032223101233223-1023310302320101-2320320032132221-0031023101132320-2230232211332312-0332131011020102)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options

<a id="canonical-3033003333232321-1322001101201210-0020223321032031-3302033013002231-1013112203123300-2302110002301322-0131032032110323-0213010222331110"></a>

Type: `"single"`. Computed.

X-Forwarded-Client-Cert header elements to be added to requests.

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

<a id="canonical-0022111030003201-2031132021213312-3021313331131313-2122100322023100-0220221303113113-2332032222123110-1000200102233320-3323100231212000"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options`

<a id="canonical-0031322011101203-0312300020230301-3002131012031120-1201310331301030-1201100312100211-0121132233302330-0100110123111101-1020333201231202"></a>

#### `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options.xfcc_header_elements` property

Type: `["list", "string"]`. Computed.

\[Enum: XFCC\_NONE|XFCC\_CERT|XFCC\_CHAIN|XFCC\_SUBJECT|XFCC\_URI|XFCC\_DNS\]
X-Forwarded-Client-Cert header elements to be added to requests. Possible values are \`XFCC\_NONE\`,
\`XFCC\_CERT\`, \`XFCC\_CHAIN\`, \`XFCC\_SUBJECT\`, \`XFCC\_URI\`, \`XFCC\_DNS\`. Defaults to
\`XFCC\_NONE\`.

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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  }
}
```

<a id="canonical-0010123223300323-3122033232120131-1033211323302101-0202012323131002-0013101320130212-1220113010330021-1032021212320311-0300031310102303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-020.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters

<a id="canonical-0103223033033133-2133130101020020-3033101100330100-3323031133022120-3113110221332132-3023032211123300-0012111230333020-3010133030032332"></a>

Type: `"single"`. Computed.

Configuration parameter for tls parameters.

Additional upstream details:

Inline TLS parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]"
}
```

<a id="canonical-3101131201103221-3023001230203323-3022023022000033-1302032213132230-1011011133200320-3011022211312300-3210202311002303-0210130230022113"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters`

- [no_mtls](data-sources--workload--reference--group-020.md#canonical-3231200111230001-0232113102222133-3103130022103033-3000222111220333-2002100110103310-2313323033112231-3312011302333313-2030332210211321): complete subsection reference.

- [tls_certificates](data-sources--workload--reference--group-020.md#canonical-2303102321200322-2223331031221031-0020113022300133-0320303032102020-2302213222211223-0230212333332330-2211002133222000-3201213130120001): complete subsection reference.

- [tls_config](data-sources--workload--reference--group-021.md#canonical-3112322103001202-3322033233311231-1102030211303311-3012000113330002-2313123200311203-2131113012330010-2213311202121332-1221333001020332): complete subsection reference.

- [use_mtls](data-sources--workload--reference--group-021.md#canonical-3300013223232120-0213310012303100-2201020132013200-0023231213231111-1000233001313111-1022001122331331-3200203330210203-2302331311311322): complete subsection reference.

<a id="canonical-3231200111230001-0232113102222133-3103130022103033-3000222111220333-2002100110103310-2313323033112231-3312011302333313-2030332210211321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.no_mtls` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-020.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-020.md#canonical-0010123223300323-3122033232120131-1033211323302101-0202012323131002-0013101320130212-1220113010330021-1032021212320311-0300031310102303)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.no_mtls

<a id="canonical-3203111113333010-2001103000233223-3212233233122122-0110033120133103-2132032232212131-3001110110303101-0232112322012101-0303230130120100"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2303102321200322-2223331031221031-0020113022300133-0320303032102020-2302213222211223-0230212333332330-2211002133222000-3201213130120001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-020.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-020.md#canonical-0010123223300323-3122033232120131-1033211323302101-0202012323131002-0013101320130212-1220113010330021-1032021212320311-0300031310102303)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates

<a id="canonical-3003313211202113-2021100011112331-2013113002121110-1230132120122202-1001333111220123-0130103022032000-2012030320023012-0222011311030003"></a>

Type: `"list"`. Computed.

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minItems": 1
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
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-1220001031213312-3020200120013200-3310021201233320-0202320301203130-3323002211212120-2301133130031221-3201120331012113-0100131230021301"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates`

<a id="canonical-0021320233012323-1220012022030003-0331031132103220-3022022033230312-3111311003020110-2001221131131231-2320330301202330-3133233100032100"></a>

#### `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.certificate_url` property

Type: `"string"`. Computed.

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

- [custom_hash_algorithms](data-sources--workload--reference--group-021.md#canonical-0232323111210030-2222131202223002-1211111033031030-3032033203332223-0312232001123030-0222030200031310-1123300111221303-0312303221022212): complete subsection reference.

<a id="canonical-0230200013123230-2303103233120212-1013322212233010-0113331221120210-3301302011202101-3112311111312030-1133323212301023-1323022323210020"></a>

<a id="canonical-0330200022230010-3230133011212210-2032223103002203-3122232233212220-3313210111023031-2100300101101301-3111110232212310-1010131101031200"></a>

#### `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.description_spec` property

Type: `"string"`. Computed.

Description. Description for the certificate.

- [disable_ocsp_stapling](data-sources--workload--reference--group-021.md#canonical-1101301113313023-1230011023121231-0220312222031231-1100013220330120-1003301213110123-2022202213001332-0313001302220030-2332211303000222): complete subsection reference.

- [private_key](data-sources--workload--reference--group-021.md#canonical-2000300020301313-0322332230111120-0301330331223200-3201132103312013-3200033102320123-1213320002213032-3123220302222201-3220313111200022): complete subsection reference.

- [use_system_defaults](data-sources--workload--reference--group-021.md#canonical-0233103130232303-0221321110230320-2111122030231332-1221212230221303-3221122020101003-2032122221233310-0132331130032110-2211011120331321): complete subsection reference.
