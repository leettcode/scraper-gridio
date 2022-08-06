# Coding
Unfinished, see video recording https://drive.google.com/file/d/1aRoefP3W2j2DmUelyiJ4bSGmjD2MB0x9/view?usp=sharing
Based on closed source personal project structure that I plan to opensource

and hexagonal architecture https://herbertograca.com/2017/11/16/explicit-architecture-01-ddd-hexagonal-onion-clean-cqrs-how-i-put-it-all-together/

and go standards https://github.com/golang-standards/project-layout

Code is taken from similar past project and altered, I've removed most stuff

Didn't manage to cover with tests but you can check for reference the code style

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
