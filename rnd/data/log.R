# Copyright 2012 Dorival de Moraes Pedroso. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

X <- seq(0, 3, 0.25)
M <- c(0, 0.5, 1)
S <- c(0.25, 0.5, 1)
Y <- matrix(ncol=4)
first <- TRUE
for (mu in M) {
    for (sig in S) {
        ypdf <- dlnorm(X, mu, sig)
        ycdf <- plnorm(X, mu, sig)
        for (i in 1:length(X)) {
            if (first) {
                Y <- rbind(c(X[i], mu, sig, ypdf[i], ycdf[i]))
                first <- FALSE
            } else {
                Y <- rbind(Y, c(X[i], mu, sig, ypdf[i], ycdf[i]))
            }
        }
    }
}
write.table(Y, "/tmp/gosl-rnd-log.dat", row.names=FALSE, col.names=c("x","mu","sig","ypdf","ycdf"), quote=FALSE)
print("file </tmp/gosl-rnd-log.dat> written")
