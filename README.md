# Problem solving

Microservice architecture brainstorming
- database to historic energy price data
- mechanism to retrieve historic price data e.g. via polling from open APIs
- scheduled service to compute most optimal vehicle charging model (presumably slow) e.g. daily
- integrate OEM API, ability to authenticate (probly involve FE, refresh auth), retrieve the status of the vehicle, command the vehicle
- define real-time charging algorithm that utilises the aforementioned model

Risks:
- complexity of algorithm. Would recomment to solve incrementally:
  1. achieve simple charge ability integrated with OEM
  2. add simple estimation whether you can achieve fully charged vehicle by morning, using charging and vehicle spec kwh etc.
  3. increase complexity by developing the efficiency model and charging algorithm (can take up to 10x time to reach this milestone, and will be ongoing effort to improve)
- missing data, open systems down
- vehicle connection issues
- keeping up with webpages and APIs updates
- availability, making sure the user has a charged car no matter what. 
  - In case the serice is to be provided globally (US/EU), different data regulations, storing local data in regions. I see no issue in syncing until no user-interaction features.
  - instance replication
- security, how do we store user credentials/tokens securely, avoid token forging (e.g. to unlock door)