# Tailscale Source Plugin

import { getLatestVersion } from "../../../../../utils/versions";
import { Badge } from "../../../../../components/Badge";

<Badge text={"Latest: " + getLatestVersion("source", `tailscale`)}/>

The CloudQuery Tailscale plugin pulls configuration out of Tailscale
resources and loads it into any supported CloudQuery destination (e.g. PostgreSQL).

## Authentication

In order to fetch information from Tailscale, `cloudquery` needs to be authenticated.
An API key is required for authentication.
Get your API key from [Tailscale Keys Settings Page](https://login.tailscale.com/admin/settings/keys).