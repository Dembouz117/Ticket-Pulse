import express from 'express';
import helmet from 'helmet';

import router from "../routes/index";


require('dotenv').config({ path: `../.env` });

const cors = require('cors');
const cookieParser = require('cookie-parser');

const app = express();
app.use(cookieParser());
// Use helmet to set security-related headers
app.use(helmet());
// Disable X-Powered-By header
app.disable('x-powered-by');

const PORT = 8082;
const allowedOrigins = ["http://localhost:3000", "http://sg1.biddlr.com", "https://sg1.biddlr.com"];
type Origin = string | undefined;
type CorsCallback = (err: Error | null, allow?: boolean) => void;


app.use(
	cors({
	  origin: function (origin: Origin, callback: CorsCallback) {
		if (!origin) return callback(null, true);
		if (allowedOrigins.indexOf(origin) === -1) {
		  const msg = 'The CORS policy for this site does not allow access from the specified Origin.';
		  return callback(new Error(msg), false);
		}
		return callback(null, true);
	  },
	  methods: 'GET,HEAD,PUT,PATCH,POST,DELETE',
	  credentials: true,
	})
  );

app.use(router);



app.listen(PORT, () => {
	console.log(`Server is running on ${PORT}`);
});
