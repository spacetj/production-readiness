# Production Readiness

What does it mean for an application to be production ready? Thankfully industry leaders a lot smarter than me have given this a lot of thought!

If you havn't come across these resources yet, it is recommended  to go through them before reading this post:

- 📹 [Building Production Ready Application][production-ready-talk] by Michael Kehoe
- 📕 [Production-Ready Microservices][production-readiness-book] by Susan J. Fowler

There are 8 main tenets of production readiness that are highlighted in these resources:

- 🔗 Stability
- 🎯 Reliability
- 🥁 Performance
- 📈 Scalability
- ✂️ Fault Tolerance
- 💥 Disaster Recovery
- 🔍 Monitoring
- 📕 Documentation

As technology evolves, several tools have emerged which helps achieve the tenets of production readiness mentioned above. This repo is part of the 4 series blog posts which covers cloud native tooling that enables application production readiness. The blogs can be found [here](https://tejasc.com/).

**DISCLAIMER** The tools mentioned in this blog are just examples of tools available, and might not be the best suited for your usecase. The blog can serve as a starting point to research the tools that is most appropriate for your ecosystem.

### Assumptions

I'll be exploring application that are:

- Running in kubernetes
- Exploring tools that are Cloud Native mostly from [CNCF Landscape][cncf-landspace]

> WARNING: The applications themselves are not production ready. The apps and tools in this repo are solely for demo and experimenting purporses.

## Stability and Reliability

Blog: [CloudNative Production Readiness Part 1](https://tejasc.com/cnpr-part1/)

Following tools have been demoed which helps build stable and reliable applications / app dev:

| Category              | Tools     |
| --------------------- | --------- |
| Application Packaging | Docker    |
|                       | Kaniko    |
|                       | Buildpack |
|                       | KO        |
| Local Testing         | Kind      |
| Manifest Packaging    | Helm      |
|                       | Tanka     |
|                       | KO        |
| CD                    | ArgoCD    |
