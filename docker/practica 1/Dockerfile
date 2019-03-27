# Use an official Python runtime as a parent image
FROM httpd

# Set the working directory to /app
WORKDIR /var/www/html

# Copy the current directory contents into the container at /app
COPY . /var/www/html

# Install any needed packages specified in requirements.txt
# RUN pip install --trusted-host pypi.python.org -r requirements.txt

# Make port 80 available to the world outside this container
EXPOSE 80

# Define environment variable
ENV NAME Server

# Run app.py when the container launches
#CMD tail -f /dev/null
CMD ["sh", "server.sh"]