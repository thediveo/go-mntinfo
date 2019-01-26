/*

Package mntinfo provides information about the currently mounted filesystems
on Linux (from the current mount namespace). This information is gathered from
the proc filesystem, in particular, from /proc/self/mountinfo, or
alternatively, from a specific PID (via /proc/[PID]/mountinfo). Just to
emphasize this point: absolutely NO /etc/fstab is used here.

Technical Details

For more background information about the mount information returned, please
see also http://man7.org/linux/man-pages/man5/proc.5.html.

Alternatives

For a multi-platform solution, please take a look at the Go gopsutil/disk
package instead (https://godoc.org/github.com/shirou/gopsutil/disk).

*/
package mntinfo
