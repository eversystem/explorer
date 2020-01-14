import { Request, Response, NextFunction } from "express";
import { veryfyJWT } from "../utils";
import { UnauthorizedError } from "../modules/error";

const verify = (req: Request, res: Response, next: NextFunction): void => {
  if (!req.headers.authorization) {
    throw new UnauthorizedError("Authorization Failed");
  }
  if (req.headers.authorization.indexOf("Bearer ") === -1) {
    throw new UnauthorizedError("Authorization Failed");
  }

  const token = req.headers.authorization.replace("Bearer ", "");

  const isValidJWT = veryfyJWT(token);

  if (!isValidJWT) {
    throw new UnauthorizedError("Authorization Failed");
  }
  console.log("token verified!!");
  next();
};

export default verify;
