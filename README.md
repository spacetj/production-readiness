# Production Readiness

What does it mean for an application to be production-ready? Thankfully, industry leaders a lot smarter than me have given this a lot of thought!

If you haven't come across these resources yet, it is recommended  to go through them before reading this post:

- 📹 [Building Production-Ready Application][production-ready-talk] by Michael Kehoe
- 📕 [Production-Ready Microservices][production-readiness-book] by Susan J. Fowler

There are eight fundamental tenets of production readiness highlighted in these resources:

- 🔗 Stability
- 🎯 Reliability
- 🥁 Performance
- 📈 Scalability
- ✂️ Fault Tolerance
- 💥 Disaster Recovery
- 🔍 Monitoring
- 📕 Documentation

As technology evolves, several tools have emerged which helps achieve the tenets of production readiness mentioned above. This repo is part of a 4 series blog post which covers cloud native tooling that enables application production readiness. The blogs can be found [here](https://tejasc.com/).

**DISCLAIMER** The tools mentioned in this blog are just examples of tools available and might not be the best suited for your use case. The blog can serve as a starting point to research the most appropriate tools for your ecosystem.

### Assumptions

I'll be exploring application that are:

- Running in Kubernetes
- Exploring tools that are Cloud Native mostly from [CNCF Landscape][cncf-landspace]

> WARNING: The applications themselves are not production ready. The apps and tools in this repo are solely for demo and experimenting.

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

<!-- Links  -->

[production-readiness-book]: https://www.amazon.com.au/Production-Ready-Microservices-Standardized-Engineering-Organization/dp/1491965975/ref=asc_df_1491965975/?tag=googleshopdsk-22&linkCode=df0&hvadid=341791741598&hvpos=&hvnetw=g&hvrand=11583469740343046994&hvpone=&hvptwo=&hvqmt=&hvdev=c&hvdvcmdl=&hvlocint=&hvlocphy=9071462&hvtargid=pla-504426002607&psc=1

[production-ready-talk]: https://www.infoq.com/presentations/production-ready-applications/

[cncf-landspace]: https://landscape.cncf.io/
