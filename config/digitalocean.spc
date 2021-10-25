connection "digitalocean" {
  plugin = "digitalocean"

  # token - (Required) Personal Access Token for your DigitalOcean account.
  # See https://www.digitalocean.com/docs/apis-clis/api/create-personal-access-token for more information.
  # Alternatively, this can be specified using environment variables ordered by precedence:
  #   - DIGITALOCEAN_TOKEN
  #   - DIGITALOCEAN_ACCESS_TOKEN
  #token = "YOUR_DIGITALOCEAN_ACCESS_TOKEN"
}
