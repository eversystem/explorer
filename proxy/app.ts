/* eslint-disable @typescript-eslint/no-unused-vars */
"use strict";
import express from "express";
import log4js from "log4js";
import http from "http";
import cors from "cors";
import cookieParser from "cookie-parser";
import axios, { AxiosResponse, AxiosError, AxiosRequestConfig } from "axios";
import awaitHandler from "./src/middlewares/awaithandler";
import { NextFunction, Request, Response } from "express";
import basicAuth from "./src/middlewares/basic-auth";
import verify from "./src/middlewares/auth";
import Status from "http-status";

log4js.configure({
  appenders: {
    out: { type: "stdout" }
  },
  categories: {
    default: { appenders: ["out"], level: "info" }
  }
});
const logger = log4js.getLogger("ATS-EXplorer-Proxy");

const host = process.env.HOST || "";
const port = process.env.PORT || 8000;

const app = express();

app.use(express.json());
app.options("*", cors());
app.use(cors());
app.use(express.urlencoded({ extended: false }));
app.use(cookieParser());

app.use(function(req: Request, res: Response, next: NextFunction) {
  if (req.originalUrl === "/health") return next();
  logger.info(" ##### New request for URL %s", req.originalUrl);
  return next();
});

// server start
const server = http.createServer(app).listen(port);
logger.info("****************** SERVER STARTED ************************");
logger.info(
  "***************  Listening on: http://%s:%s  ******************",
  host,
  port
);
server.timeout = 240000;

// health check
app.get(
  "/health",
  awaitHandler(async (req: Request, res: Response) => {
    res.sendStatus(200);
  })
);

app.post("/api/login", basicAuth);

// Auth middleware
app.use("/api/verify", verify);

app.use("/api", verify);

// Proxy api
app.use(
  "/api",
  awaitHandler(async (req: Request, res: Response) => {
    const { url, method, body } = req as {
      url: string;
      method: "GET" | "POST";
      body: any;
    };

    if (!url || !method) {
      res.status(Status.BAD_REQUEST).send("Bad Request");
    }

    if (method != "GET" && method != "POST") {
      res.status(Status.METHOD_NOT_ALLOWED).send("Method not allowed");
    }

    const result: AxiosResponse | void = await axios({
      method,
      url,
      data: body
    }).catch((err: AxiosError) => {
      res.status(500).send("Internal Server Error");
    });

    if (!result) {
      return res.send();
    }

    res.send(result.data);
  })
);

/** **********************************************************************************
 * Error handler
 ************************************************************************************/
app.use(function(error: any, req: Request, res: Response, next: NextFunction) {
  console.log(error);
  if (error.status) {
    res.status(error.status);
  } else {
    res.status(500);
  }
  res.json({ message: error.message }).send();
});
