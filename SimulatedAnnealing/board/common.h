#ifndef EMSA_COMMON_H
#define EMSA_COMMON_H
/*
 *  Simulated Annealing Symbolics and Function Prototypes
 *
 *  ./software/ch2/emsa/common.h
 *
 *  mtj@cogitollc.com
 *
 */

#include <stdlib.h>

static int MAX_LENGTH;

typedef int* solutionType;

typedef struct {
  solutionType solution;
  float energy;
} memberType;

/* Annealing Schedule */
static double INITIAL_TEMPERATURE;
static double FINAL_TEMPERATURE;
static double ALPHA;
static int STEPS_PER_CHANGE;

/* Return a random number between 0 and 1 */
#define getSRand()	((float)rand() / (float)RAND_MAX)

/* Return an integer from 0..(x-1) */
#define getRand(x)	(int)(rand() % x)

#endif //EMSA_COMMON_H/*
/*
 *  Copyright (c) 2003 Charles River Media.  All rights reserved.
 * 
 *  Redistribution and use in source and binary forms, with or 
 *  without modification, is hereby granted without fee provided 
 *  that the following conditions are met:
 * 
 *    1.  Redistributions of source code must retain the above 
 *        copyright notice, this list of conditions and the 
 *        following disclaimer.
 *    2.  Redistributions in binary form must reproduce the above
 *        copyright notice, this list of conditions and the 
 *        following disclaimer in the documentation and/or other 
 *        materials provided with the distribution.
 *    3.  Neither the name of Charles River Media nor the names of 
 *        its contributors may be used to endorse or promote 
 *        products derived from this software without specific 
 *        prior written permission.
 * 
 * THIS SOFTWARE IS PROVIDED BY CHARLES RIVER MEDIA AND CONTRIBUTORS 
 * 'AS IS' AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT 
 * LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTIBILITY AND FITNESS 
 * FOR A PARTICULAR PURPOSE ARE DISCLAIMED.  IN NO EVENT SHALL CHARLES
 * RIVER MEDIA OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, 
 * INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, 
 * BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS 
 * OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED 
 * AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT 
 * LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN 
 * ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE 
 * POSSIBILITY OF SUCH DAMAGE.
 */

