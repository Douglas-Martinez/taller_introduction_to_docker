FROM node:16-alpine3.13 as builder
WORKDIR /app
COPY . .
RUN npm install 
CMD ["npm","start"]

# Para Produccion
# RUN npm run build
# RUN npm install -g serve 
# CMD ["serve","-s","build"]