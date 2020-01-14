// Basic認証
import "../config/env.ts";
import auth from "basic-auth";
import { NextFunction, Request, Response } from "express";
import { UnauthorizedError } from "../modules/error";
import { generateJWT } from "../utils";
import env from "../config/env";

const admins = {
  [env.BASICNAME]: { password: env.BASICPASS }
};

const basicAuth = (
  req: Request,
  res: Response,
  next: NextFunction
): void | Response => {
  console.log(req.headers.authorization);
  if (req.headers.authorization?.indexOf("Bearer") != -1) {
    const token = req.headers.authorization?.replace("Bearer ", "");
    // Bearer指定の場合はトークン認証に通す
    if (token) {
      return next();
    }
  }

  const user = auth(req);
  console.log(user);
  if (!user || !admins[user.name] || admins[user.name].password !== user.pass) {
    throw new UnauthorizedError("Authorization Failed");
  }

  const jwt = generateJWT();

  return res.json({ jwt }).send();
};

export default basicAuth;
